package format

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

// ThreagileOutput represents the Threagile schema format
type ThreagileOutput struct {
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Assets      map[string]ThreagileAsset `json:"assets"`
}

type ThreagileAsset struct {
	Name       string   `json:"name"`
	DependsOn  []string `json:"depends-on"`
	RiskLevels []string `json:"risk-levels"`
}

func Threagile(config quickrisk.Config) {
	threagile := ThreagileOutput{
		Title:       "Generated Threat Model",
		Description: "This threat model is generated from YAML configuration",
		Assets:      make(map[string]ThreagileAsset),
	}

	for componentName, component := range config.Components {
		asset := ThreagileAsset{
			Name:      componentName,
			DependsOn: component.Deps,
		}

		for riskName, risk := range component.Risks {
			if risk != nil {
				severity := "low"
				if risk.Score >= 3 {
					severity = "high"
				} else if risk.Score == 2 {
					severity = "medium"
				}
				asset.RiskLevels = append(asset.RiskLevels, fmt.Sprintf("%s: %s", riskName, severity))
			}
		}

		threagile.Assets[componentName] = asset
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(threagile); err != nil {
		log.Fatalf("Failed to write Threagile output: %v", err)
	}
}
