package quickrisk

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadConfigs(paths []string) (Config, error) {
	c := Config{
		Components: map[string]*Component{},
	}

	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
				return nil
			}

			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("Skipping %s: %v", path, err)
				return nil
			}

			var this Config
			err = yaml.Unmarshal(data, &this)
			if err != nil {
				return fmt.Errorf("unmarshal %s: %w", path, err)
			}

			// TODO: consider edge cases when merging components
			for k, v := range this.Components {
				if c.Components[k] != nil {
					return fmt.Errorf("%s: duplicate key: %s", path, k)
				}

				if v == nil {
					continue
				}
				for _, r := range v.Risks {
					if r == nil {
						r = &Risk{}
					}
					if r.UnmitigatedScore == 0 {
						r.UnmitigatedScore = (r.Impact - 2) + (r.Likelihood - 2)
					}
					if r.Score == 0 {
						r.Score = r.UnmitigatedScore
						for _, m := range r.Mitigations {
							r.Score += m
						}
					}
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
