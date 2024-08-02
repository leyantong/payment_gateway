package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	BankSimulatorURL string `mapstructure:"bank_simulator_url"`
	Port             string `mapstructure:"port"`
	DatabaseURL      string `mapstructure:"database_url"`
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
