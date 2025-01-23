package format

import (
	"fmt"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func DOT(config quickrisk.Config) {
	fmt.Println("digraph Components {")

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
		fmt.Printf("\t\"%s\" [color=%s];\n", componentName, color)

		// Print dependencies as edges
		for _, dep := range component.Deps {
			fmt.Printf("\t\"%s\" -> \"%s\";\n", componentName, dep)
		}
	}

	fmt.Println("}")
}
