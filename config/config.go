package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// variables

var DbConfig DbConfigStruct = initDbConfig()

var ServerConfig ServerConfigStruct = initServerConfig()

// types

type DbConfigStruct struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

type ServerConfigStruct struct {
	Port               string
	JWT_SECRET         string
	Jwt_REFRESH_SECRET string
}

// funcs

func initDbConfig() DbConfigStruct {
	return DbConfigStruct{
		DB_HOST: GetEnv("DB_HOST"),
		DB_PORT: GetEnv("DB_PORT"),
		DB_USER: GetEnv("DB_USER"),
		DB_PASS: GetEnv("DB_PASS"),
		DB_NAME: GetEnv("DB_NAME"),
	}
}

func initServerConfig() ServerConfigStruct {
	return ServerConfigStruct{
		Port:               GetEnv("PORT"),
		JWT_SECRET:         GetEnv("JWT_SECRET"),
		Jwt_REFRESH_SECRET: GetEnv("Jwt_REFRESH_SECRET"),
	}
}

func GetEnv(key string) string {
	// connect to env
	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load envs")
	}
	// get env
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("could not found %s env", key)
	}
	return value
}

func GetEnvAsInt(key string) int64 {
	// connect to env
	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load envs")
	}
	// get env
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("could not found %s env", key)
	}
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatalf("could not parse to int %s env", key)
	}
	return parsedValue
}
