package quickrisk

// Risk represents an individual risk item
type Risk struct {
	Impact      int            `yaml:"impact"`
	Likelihood  int            `yaml:"likelihood"`
	Mitigations map[string]int `yaml:"mitigations"`
	// Generally a calculated field
	Score int `yaml:"score"`
	// Generally a calculated field
	UnmitigatedScore int `yaml:"unmitigated_score"`
}

// Component represents a single component's configuration
type Component struct {
	Risks  map[string]*Risk `yaml:"risks"`
	Deps   []string         `yaml:"deps"`
	Trusts []string         `yaml:"trusts"`
	Has    []string         `yaml:"has"`
}

// Config represents the overall YAML structure
type Config struct {
	Components map[string]*Component `yaml:"components"`
	Default    *Risk                 `yaml:"default"`
}
