package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int
	MongoURI string
	RiotAPIKey string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		Port: viper.GetInt("PORT"),
		MongoURI: viper.GetString("MONGO_URI"),
		RiotAPIKey: viper.GetString("RIOT_API_KEY"),
	}
	return cfg, nil
}