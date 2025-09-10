package config

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
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

	port , err := strconv.ParseInt(httpPort,10,64)

	if err != nil {
		fmt.Println("Couldn't convert into integer")
	}

	configurations = Config {
		Version: version,
		ServiceName: serviceName,
		HttpPort: int(port),
	}

}

func GetConfig() Config {
	loadConfig()
	return configurations
}
