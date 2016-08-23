package app

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/go-ozzo/ozzo-validation"
)

var Config appConfig

type appConfig struct {
	ErrorFile          string `mapstructure:"error_file"`
	ServerPort         int    `mapstructure:"server_port"`
	DSN                string `mapstructure:"dsn"`
	JWTSigningKey      string `mapstructure:"jwt_signing_key"`
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
}

func (config *appConfig) Validate() error {
	return validation.StructRules{}.
		Add("ServerPort", validation.Required).
		Add("DSN", validation.Required).
		Add("JWTSigningKey", validation.Required).
		Add("JWTVerificationKey", validation.Required).
		Validate(config)
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault("error_file", "config/errors.yaml")
	v.SetDefault("server_port", 8080)
	v.SetEnvPrefix("restful")
	v.SetConfigType("yaml")
	v.SetConfigName("app")
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
