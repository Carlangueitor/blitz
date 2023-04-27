package blitz

type Config struct {
	Port string `mapstructure:"PORT"`
}

type ConfigLoader interface {
	Load() Config
}
