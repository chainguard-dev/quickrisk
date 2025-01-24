package format

import (
	"fmt"
	"io"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func DOT(w io.Writer, config quickrisk.Config) {
	fmt.Fprintln(w, "digraph Components {")

	for componentName, component := range config.Components {
		// Determine node color based on risk
		highRisk := false
		for _, risk := range component.Risks {
			if risk != nil {
				if risk.Score >= 3 {
					highRisk = true
					break
				}
			}
		}
		color := "black"
		if highRisk {
			color = "red"
		}
		fmt.Fprintf(w, "\t\"%s\" [color=%s];\n", componentName, color)

		// Print dependencies as edges
		for _, dep := range component.Deps {
			fmt.Fprintf(w, "\t\"%s\" -> \"%s\";\n", componentName, dep)
		}
	}

	fmt.Fprintln(w, "}")
}
