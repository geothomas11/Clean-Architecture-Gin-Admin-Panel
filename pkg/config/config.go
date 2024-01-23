package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBName     string `mapstructure: "DBNAME"`
	DBPassword string `mapstructure:"PASSWORD"`
	DBHost     string `mapstructure:"HOST"`
	DBPort     string `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}
