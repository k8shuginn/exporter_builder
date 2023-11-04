package config

type Config struct {
	Name       string   `mapstructure:"name"`
	Module     string   `mapstructure:"module"`
	Collectors []string `mapstructure:"collectors"`
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{
		Name:   "diy_exporter",
		Module: "diy_exporter",
		Collectors: []string{
			"sample1",
			"sample2",
		},
	}
}
