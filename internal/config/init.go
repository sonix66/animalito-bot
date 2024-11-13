package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig(configFile string) (*Config, error) {
	v := viper.New()
	ext := strings.TrimLeft(filepath.Ext(configFile), ".")
	v.SetConfigFile(configFile)
	v.SetConfigType(ext)

	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("v.ReadInConfig: %w", err)
	}

	for _, k := range v.AllKeys() {
		value := v.GetString(k)
		if value == "" {
			continue
		}
		v.Set(k, os.ExpandEnv(value))
	}

	cfg := new(Config)
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("v.Unmarshal: %w", err)
	}

	return cfg, nil
}
