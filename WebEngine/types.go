package WebEngine

type ServerConfig struct {
	Location string `yaml:"location"`
	Port     uint `yaml:"port"`
}

type TemplateConfig struct {
	HtmlRoot string `yaml:"htmlRoot"`
	StaticRoot string `yaml:"staticRoot"`
}