package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func MustInit(configFile string) *Config {
	v := viper.New()
	ext := strings.TrimLeft(filepath.Ext(configFile), ".")
	v.SetConfigFile(configFile)
	v.SetConfigType(ext)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("v.ReadInConfig: %w", err))
	}

	for _, k := range v.AllKeys() {
		value := v.GetString(k)
		if value == "" {
			continue
		}
		v.Set(k, os.ExpandEnv(value))
	}

	var cfg *Config

	err = v.Unmarshal(cfg)
	if err != nil {
		panic(fmt.Errorf("v.Unmarshal: %w", err))
	}

	return cfg
}
