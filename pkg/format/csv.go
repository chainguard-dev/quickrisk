package format

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func CSV(w io.Writer, config quickrisk.Config) {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write CSV header
	err := writer.Write([]string{"Mitigated Risk Score", "Zone", "Component", "Risk Name", "Has", "Unmitigated Risk Score", "Mitigations", "Impact", "Likelihood"})
	if err != nil {
		log.Fatalf("Failed to write CSV header: %v", err)
	}

	// Write CSV rows
	for name, c := range config.Components {
		for riskName, r := range c.Risks {
			if r == nil {
				continue
			}
			ms := []string{}
			for m := range r.Mitigations {
				ms = append(ms, m)
			}
			sort.Strings(ms)

			hs := []string{}
			for _, h := range c.Has {
				hs = append(hs, h)
			}
			sort.Strings(hs)

			row := []string{fmt.Sprintf("%d", r.Score), c.Zone, name, riskName, strings.Join(hs, ", "), fmt.Sprintf("%d", r.UnmitigatedScore), strings.Join(ms, ", "), fmt.Sprintf("%d", r.Impact), fmt.Sprintf("%d", r.Likelihood)}
			err = writer.Write(row)
			if err != nil {
				log.Fatalf("Failed to write CSV row: %v", err)
			}
		}
	}
}
