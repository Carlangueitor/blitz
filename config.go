package blitz

//go:generate mockgen -destination=./mocks/mock_config.go -package=mocks github.com/carlangueitor/blitz ConfigLoader

type Config struct {
	Port int `mapstructure:"PORT"`
}

type ConfigLoader interface {
	Load() (*Config, error)
}
