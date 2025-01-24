package format

import (
	"fmt"
	"io"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func DOT(w io.Writer, config quickrisk.Config) {
	fmt.Fprintln(w, "digraph Components {")

	// Map to group components by zones
	zones := make(map[string][]string)

	for componentName, component := range config.Components {
		// Determine node color based on risk
		highRisk := false
		for _, risk := range component.Risks {
			if risk != nil && risk.Score >= 3 {
				highRisk = true
				break
			}
		}
		color := "black"
		if highRisk {
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

		// Print trust relationships as blue dotted edges
		for _, t := range component.Trusts {
			fmt.Fprintf(w, "\t\"%s\" <- \"%s\" [style=dotted, color=blue];\n", t, componentName)
		}
	}

	// Print zones as subgraphs
	for zoneName, components := range zones {
		fmt.Fprintf(w, "\tsubgraph \"cluster_%s\" {\n", zoneName)
		fmt.Fprintln(w, "\t\tstyle=dashed;")
		fmt.Fprintln(w, "\t\tcolor=gray50;")
		fmt.Fprintf(w, "\t\tlabel=\"%s\";\n", zoneName)
		for _, comp := range components {
			fmt.Fprintln(w, comp)
		}
		fmt.Fprintln(w, "\t}")
	}

	fmt.Fprintln(w, "}")
}
