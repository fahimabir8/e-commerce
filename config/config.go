package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DbConfig struct {
	User          string
	Password      string
	Host          string
	Port          int
	DBName        string
	EnableSSLMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DbConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env files: ", err)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("Couldn't convert into integer")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT secret key requied!!")
		os.Exit(1)
	}

	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		fmt.Println("DB user is required!!")
		os.Exit(1)
	}

	dbpass := os.Getenv("DB_PASSWORD")
	if dbpass == "" {
		fmt.Println("DB password is required!!")
		os.Exit(1)
	}

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		fmt.Println("DB host is required!!")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB port is required")
		os.Exit(1)
	}

	DBport, err := strconv.ParseInt(dbPort, 10, 64)

	if err != nil {
		fmt.Println("Couldn't convert into integer")
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		fmt.Println("DB name is required")
		os.Exit(1)
	}

	is_enable_ssl := os.Getenv("DB_ENABLE_SSL_MODE")

	sslMode, err := strconv.ParseBool(is_enable_ssl)
	if err != nil {
		fmt.Println("Couldn't parse to boolean",err)
		os.Exit(1)
	}

	DBConfig := &DbConfig{
		User:          dbuser,
		Password:      dbpass,
		Host:          dbhost,
		Port:          int(DBport),
		DBName:        dbname,
		EnableSSLMode: sslMode,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           DBConfig,
	}

}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
