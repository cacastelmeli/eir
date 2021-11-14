package config

type TemplateHook struct {
	// Command name
	// Only supported for: `init` and `type` commands
	CommandName string `yaml:"cmd"`

	// Pre hook path
	PrePath string `yaml:"pre"`

	// Post hook path
	PostPath string `yaml:"post"`
}

type TemplateConfig struct {
	Hooks TemplateHook `yaml:"hooks"`
}
