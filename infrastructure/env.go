package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	ENVIRONMENT string `mapstructure:"ENVIRONMENT"`
	LogOutput   string `mapstructure:"LOG_OUTPUT"`
	DbUsername  string `mapstructure:"DB_USERNAME"`
	DbPassword  string `mapstructure:"DB_PASSWORD"`
	DbHost      string `mapstructure:"DB_HOST"`
	DbPort      string `mapstructure:"DB_PORT"`
	DbName      string `mapstructure:"DB_NAME"`
	DbType      string `mapstructure:"DB_TYPE"`
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read configuration")
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}
	log.Printf("%#v \n", env)
	return env
}
