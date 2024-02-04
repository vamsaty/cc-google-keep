package config

import (
	"encoding/json"
	"io"
	"os"
)

const kAppConfigFile = "configs/app_config.json"

// LoadAppConfigFrom loads the app config from the given source.
func LoadAppConfigFrom(source string) AppConfig {
	var err error
	var file io.ReadCloser

	if source == "" {
		source = kAppConfigFile
	}

	if file, err = os.Open(source); err != nil {
		panic(err)
	}
	defer file.Close()

	var config AppConfig
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		panic(err)
	}
	return config
}

// LoadAppConfig loads the app config from the default source.
func LoadAppConfig() AppConfig {
	return LoadAppConfigFrom("")
}
