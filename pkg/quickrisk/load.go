package quickrisk

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Validate(c Config) []error {
	errs := []error{}

	found := map[string]*Component{}
	for k, v := range c.Components {
		found[k] = v
	}

	foundZone := map[string]bool{}
	for _, v := range c.Components {
		foundZone[v.Zone] = true
	}

	for k, v := range c.Components {
		for _, d := range v.Deps {
			if found[d] == nil {
				errs = append(errs, fmt.Errorf("component %q references unknown dependency %q", k, d))
			}
		}
		for _, d := range v.ZoneDeps {
			if !foundZone[d] {
				errs = append(errs, fmt.Errorf("component %q references unknown zone dependency %q", k, d))
			}
		}
		for _, d := range v.Trusts {
			if found[d] == nil {
				errs = append(errs, fmt.Errorf("component %q trusts unknown dependency %q", k, d))
			}
		}

	}
	return errs
}

func LoadConfigs(paths []string) (Config, error) {
	c := Config{
		Components: map[string]*Component{},
	}

	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || (filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml") {
				return nil
			}

			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("Skipping %s: %v", path, err)
				return nil
			}

			var this Config
			if err := yaml.Unmarshal(data, &this); err != nil {
				return fmt.Errorf("unmarshal %s: %w", path, err)
			}

			// Merge components with respect to defaults
			for k, v := range this.Components {
				if c.Components[k] != nil {
					return fmt.Errorf("%s: duplicate key: %s", path, k)
				}
				if v == nil {
					v = &Component{}
				}

				// Apply default component values
				applyDefaultComponentValues(v, &this.Defaults.Component)

				// Apply default risk values
				for rk, r := range v.Risks {
					if r == nil {
						r = &Risk{}
					}
					applyDefaultRiskValues(r, this.Defaults.Risk)
					v.Risks[rk] = r
				}

				c.Components[k] = v
			}
			return nil
		})
		if err != nil {
			return c, err
		}
	}

	return c, nil
}

// Helper function to apply default component values
func applyDefaultComponentValues(comp *Component, defaultComp *Component) {
	if comp.Zone == "" && defaultComp.Zone != "" {
		comp.Zone = defaultComp.Zone
	}
	if len(comp.ZoneDeps) == 0 {
		comp.ZoneDeps = append(comp.ZoneDeps, defaultComp.ZoneDeps...)
	}

	if len(comp.Deps) == 0 {
		comp.Deps = append(comp.Deps, defaultComp.Deps...)
	}
	if len(comp.Trusts) == 0 {
		comp.Trusts = append(comp.Trusts, defaultComp.Trusts...)
	}
	if len(comp.Has) == 0 {
		comp.Has = append(comp.Has, defaultComp.Has...)
	}
}

// Helper function to apply default risk values
func applyDefaultRiskValues(risk *Risk, defaultRisk *Risk) {
	if defaultRisk == nil {
		return
	}
	if risk.Impact == 0 {
		risk.Impact = defaultRisk.Impact
	}
	if risk.Likelihood == 0 {
		risk.Likelihood = defaultRisk.Likelihood
	}
	if len(risk.Mitigations) == 0 {
		risk.Mitigations = map[string]int{}
		for k, v := range defaultRisk.Mitigations {
			risk.Mitigations[k] = v
		}
	}
	if risk.UnmitigatedScore == 0 {
		risk.UnmitigatedScore = (risk.Impact - 2) + (risk.Likelihood - 2)
	}
	if risk.Score == 0 {
		risk.Score = risk.UnmitigatedScore
		for _, m := range risk.Mitigations {
			risk.Score = risk.Score + (int(math.Abs(float64(m))) * -1)
		}
	}
}
