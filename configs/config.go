package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_DATABASE"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("⚠️ .env file not found, relying on environment variables", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalln("❌ Unable to decode into struct", err)
	}

	return config
}
