// Package config provides configuration-related types and functions.
package config

import (
	"github.com/spf13/viper"
	"github.com/verless/verless/model"
)

// Config represents the user configuration stored in verless.yml.
type Config struct {
	Site struct {
		Meta model.Meta
		Nav  struct {
			Items []struct {
				Label  string
				Target string
			}
			Override bool
		}
		Footer struct {
			Items []struct {
				Label  string
				Target string
			}
			Override bool
		}
	}
}

// FromFile looks for a YAML, TOML oder JSON file with the given
// name in the provided path and converts it to a Config instance.
func FromFile(path, filename string) (Config, error) {
	viper.AddConfigPath(path)
	// Set the filename without extension to allow all supported formats.
	viper.SetConfigName(filename)

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
