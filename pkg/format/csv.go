package format

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func CSV(config quickrisk.Config) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// Write CSV header
	err := writer.Write([]string{"Component", "Risk Name", "Risk Score"})
	if err != nil {
		log.Fatalf("Failed to write CSV header: %v", err)
	}

	// Write CSV rows
	for componentName, component := range config.Components {
		for riskName, risk := range component.Risks {
			if risk != nil {
				row := []string{componentName, riskName, fmt.Sprintf("%d", risk.Score)}
				err = writer.Write(row)
				if err != nil {
					log.Fatalf("Failed to write CSV row: %v", err)
				}
			}
		}
	}
}
