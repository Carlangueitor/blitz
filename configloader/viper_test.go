package configloader

import (
	"reflect"
	"testing"

	"github.com/carlangueitor/blitz"
)

func TestViperConfigLoaderDefaultValues(t *testing.T) {
	loader := ViperConfigLoader{}
	config, err := loader.Load()

	if err != nil {
		t.Errorf("Failed to load config %s", err)
	}

	if config.Port != defaultPort {
		t.Error("Default port Value was not the expected")
	}
}

func TestViperConfigLoaderLoadFromEnv(t *testing.T) {
	expectedConfig := &blitz.Config{
		Port: 5678,
	}

	t.Setenv("BLITZ_PORT", "5678")
	loader := ViperConfigLoader{}
	config, err := loader.Load()

	if err != nil {
		t.Errorf("Failed to load config %s", err)
	}

	if !reflect.DeepEqual(expectedConfig, config) {
		t.Errorf("Loaded config was diferent than the expected one: expected=%v actual=%v", expectedConfig, config)
	}
}
