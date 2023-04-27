package blitz

type Config struct {
	Port int `mapstructure:"PORT"`
}

type ConfigLoader interface {
	Load() (*Config, error)
}
