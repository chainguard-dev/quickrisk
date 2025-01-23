package format

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

// OTMOutput represents the Open Threat Modeling Format
type OTMOutput struct {
	OTMVersion string         `json:"otmVersion"`
	Project    OTMProject     `json:"project"`
	Components []OTMComponent `json:"components"`
}

type OTMProject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OTMComponent struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Risks        []OTMRisk `json:"risks"`
	Dependencies []string  `json:"dependencies"`
}

type OTMRisk struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Severity string `json:"severity"`
}

func OTM(config quickrisk.Config) {
	otm := OTMOutput{
		OTMVersion: "1.1",
		Project: OTMProject{
			Name:        "Example Project",
			Description: "Generated from YAML configuration",
		},
	}

	for componentName, component := range config.Components {
		otmComponent := OTMComponent{
			Id:           componentName,
			Name:         componentName,
			Type:         "Generic Component",
			Dependencies: component.Deps,
		}

		for riskName, risk := range component.Risks {
			if risk != nil {
				severity := "Low"
				if risk.Score >= 3 {
					severity = "High"
				} else if risk.Score == 2 {
					severity = "Medium"
				}
				otmRisk := OTMRisk{
					Id:       fmt.Sprintf("%s_%s", componentName, riskName),
					Name:     riskName,
					Severity: severity,
				}
				otmComponent.Risks = append(otmComponent.Risks, otmRisk)
			}
		}

		otm.Components = append(otm.Components, otmComponent)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(otm); err != nil {
		log.Fatalf("Failed to write OTM output: %v", err)
	}
}
