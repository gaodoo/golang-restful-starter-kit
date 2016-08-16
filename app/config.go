package app

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
}

func LoadConfig(paths ...string) (Config, error) {
	v := viper.New()
	v.SetDefault("error_file", "config/errors.yaml")
	v.SetDefault("server_port", "8080")

	v.SetEnvPrefix("restful")
	v.BindEnv("dsn")
	v.BindEnv("server_port")

	v.SetConfigType("yaml")
	v.SetConfigName("app")
	for _, path := range paths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	return v, nil
}
