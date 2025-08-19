package config

type InferenceConfig struct {
	APIVersion string      `yaml:"apiVersion"`
	Debug      bool        `yaml:"debug"`
	Runtime    string      `yaml:"runtime"`
	Backends   []string    `yaml:"backends"`
	Models     []Model     `yaml:"models"`
	Config     string      `yaml:"config"`
	Devices    []CDIDevice `yaml:"devices,omitempty"`
}

type CDIDevice struct {
	Name     string `yaml:"name"`
	Required bool   `yaml:"required"`
	Class    string `yaml:"class,omitempty"`
}

type Model struct {
	Name            string           `yaml:"name"`
	Source          string           `yaml:"source"`
	SHA256          string           `yaml:"sha256"`
	PromptTemplates []PromptTemplate `yaml:"promptTemplates"`
}

type PromptTemplate struct {
	Name     string `yaml:"name"`
	Template string `yaml:"template"`
}
