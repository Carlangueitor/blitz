package configloader

import (
	"github.com/spf13/viper"

	"github.com/carlangueitor/blitz"
)

const defaultPort = 4000

type ViperConfigLoader struct {
}

func (loader *ViperConfigLoader) Load() (*blitz.Config, error) {
	viper.SetDefault("port", defaultPort)
	viper.SetEnvPrefix("blitz")
	viper.AutomaticEnv()

	config := blitz.Config{}

	err := viper.Unmarshal(&config)
	return &config, err
}
