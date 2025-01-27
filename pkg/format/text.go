package format

import (
	"fmt"
	"io"

	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func Text(w io.Writer, config quickrisk.Config) {
	// Calculate risk scores and print the parsed data
	for name, c := range config.Components {
		fmt.Fprintf(w, "[%s]\n", name)

		if c == nil {
			fmt.Println("  No data.")
			continue
		}

		if len(c.Has) > 0 {
			fmt.Fprintln(w, "  has:")
			for _, v := range c.Has {
				fmt.Fprintf(w, "    -  %s\n", v)
			}
		}

		if len(c.Deps) > 0 {
			fmt.Fprintln(w, "  dependencies:")
			for _, dep := range c.Deps {
				fmt.Fprintf(w, "    - %s\n", dep)
			}
		}

		if c.Zone != "" {
			fmt.Fprintf(w, "  zone: %s\n", c.Zone)
		}

		if len(c.ZoneDeps) > 0 {
			fmt.Fprintln(w, "  zone dependencies:")
			for _, dep := range c.ZoneDeps {
				fmt.Fprintf(w, "    - %s\n", dep)
			}
		}

		fmt.Fprintln(w, "  risks:")
		for riskName, r := range c.Risks {
			fmt.Fprintf(w, "    %s:\n", riskName)
			if r != nil {
				fmt.Fprintf(w, "      impact: %d\n", r.Impact)
				fmt.Fprintf(w, "      likelihood: %d\n", r.Likelihood)
				fmt.Fprintf(w, "      risk score: %d\n", r.Score)
				if len(r.Mitigations) > 0 {
					fmt.Fprintln(w, "      mitigations:")
					for k, v := range r.Mitigations {
						fmt.Fprintf(w, "        %s: %d\n", k, v)
					}
				}
			}
		}

		fmt.Fprintln(w, "")
	}
}
