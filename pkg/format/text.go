package format

import (
	"fmt"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func Text(config quickrisk.Config) {
	// Calculate risk scores and print the parsed data
	for componentName, component := range config.Components {
		fmt.Printf("Component: %s\n", componentName)

		if component == nil {
			fmt.Println("  No risks.")
			continue
		}

		// Print risks
		fmt.Println("  Risks:")
		for riskName, risk := range component.Risks {
			fmt.Printf("    %s:\n", riskName)
			if risk != nil {
				fmt.Printf("      Impact: %d\n", risk.Impact)
				fmt.Printf("      Likelihood: %d\n", risk.Likelihood)
				fmt.Printf("      Risk Score: %d\n", risk.Score)
				if len(risk.Mitigations) > 0 {
					fmt.Println("      Mitigations:")
					for mitigation, value := range risk.Mitigations {
						fmt.Printf("        %s: %d\n", mitigation, value)
					}
				}
			}
		}

		// Print dependencies
		fmt.Println("  Dependencies:")
		for _, dep := range component.Deps {
			fmt.Printf("    - %s\n", dep)
		}
	}
}
