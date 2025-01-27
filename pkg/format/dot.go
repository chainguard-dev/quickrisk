package format

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

var (
	nonWordRe = regexp.MustCompile(`\W+`)
)

func DOT(w io.Writer, config quickrisk.Config) {
	fmt.Fprintln(w, "digraph Components {")
	fmt.Fprintln(w, "  compound=true;")

	fmt.Fprintln(w, "  graph [fontsize=10 fontname=\"Verdana\" compound=true];")
	fmt.Fprintln(w, "  node [shape=record fontsize=10 fontname=\"Verdana\"];")

	// Map to group components by zones
	zones := make(map[string][]string)

	for componentName, component := range config.Components {
		// Determine node color based on risk
		highRisk := false
		critRisk := false

		for _, risk := range component.Risks {
			if risk != nil && risk.Score >= 4 {
				critRisk = true
				break
			}
			if risk != nil && risk.Score >= 3 {
				highRisk = true
				break
			}
		}
		color := "black"
		if highRisk {
			color = "yellow"
		}
		if critRisk {
			color = "red"
		}

		// Group components into zones
		if component.Zone != "" {
			zones[component.Zone] = append(zones[component.Zone], fmt.Sprintf("\t\"%s\" [color=%s];", componentName, color))
		} else {
			// Components without a zone
			fmt.Fprintf(w, "\t\"%s\" [color=%s];\n", componentName, color)
		}

		// Print dependencies as edges
		for _, dep := range component.Deps {
			fmt.Fprintf(w, "\t\"%s\" -> \"%s\";\n", componentName, dep)
		}

		for _, dep := range component.ZoneDeps {
			// TODO: fix hardcoding
			target := "auth.google.com"
			fmt.Fprintf(w, "\t\"%s\" -> \"%s\" [ltail=\"%s\" lhead=\"%s\"];\n", componentName, target, zoneGraphCluster(component.Zone), zoneGraphCluster(dep))
		}

		// Print trust relationships as blue dotted edges
		// for _, t := range component.Trusts {
		// 	fmt.Fprintf(w, "\t\"%s\" -> \"%s\" [style=dotted, color=blue];\n", componentName, t)
		// }
	}

	// Print zones as subgraphs
	for zone, components := range zones {
		fmt.Fprintf(w, "\tsubgraph \"%s\" {\n", zoneGraphCluster(zone))
		// 		fmt.Fprintln(w, "\t\tnode [style=filled];")
		fmt.Fprintf(w, "\t\tlabel=\"%s\";\n", zone)
		for _, comp := range components {
			fmt.Fprintf(w, "\t%s\n", comp)
		}
		fmt.Fprintf(w, "\t\tcolor=blue;\n")
		fmt.Fprintln(w, "\t}")
	}

	fmt.Fprintln(w, "}")
}

func zoneGraphCluster(s string) string {
	s = strings.ToLower(nonWordRe.ReplaceAllString(s, "_"))
	return fmt.Sprintf("cluster_%s", s)
}
