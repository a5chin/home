package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HOSTNAME    string `env:"hostname"`
	PORT        string `env:"port"`
	DB_HOSTNAME string `env:"dbhostname"`
	DB_USER     string `env:"dbuser"`
	DB_PWD      string `env:"dbpwd"`
	DB_NAME     string `env:"dbname"`
	DB_TCPHOST  string `env:"dbtcphost"`
	DB_PORT     string `env:"dbport"`
}

func ExistEnvFile() bool {
	_, err := os.Stat("./config/.env")
	return err == nil
}

func Load() *Config {
	var config = &Config{}
	if ExistEnvFile() {
		err := godotenv.Load("./config/.env")
		if err != nil {
			log.Println(err)
		}
		config = &Config{
			HOSTNAME:    os.Getenv("HOSTNAME"),
			PORT:        os.Getenv("PORT"),
			DB_HOSTNAME: os.Getenv("DB_HOSTNAME"),
			DB_USER:     os.Getenv("DB_USER"),
			DB_PWD:      os.Getenv("DB_PWD"),
			DB_NAME:     os.Getenv("DB_NAME"),
			DB_PORT:     os.Getenv("DB_PORT"),
		}
	} else {
		config = &Config{
			HOSTNAME:   os.Getenv("HOSTNAME"),
			PORT:       os.Getenv("PORT"),
			DB_USER:    os.Getenv("DB_USER"),
			DB_PWD:     os.Getenv("DB_PWD"),
			DB_NAME:    os.Getenv("DB_NAME"),
			DB_TCPHOST: os.Getenv("DB_TCPHOST"),
			DB_PORT:    os.Getenv("DB_PORT")}
	}
	return config
}
