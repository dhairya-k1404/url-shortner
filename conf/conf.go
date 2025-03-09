package conf

import (
	"log"
	"os"
)

var (
	BASE_URL string

	POSTGRES_HOST     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_PORT     string

	REDIS_URL string
)

func init() {
	log.Println("conf: into init func")

	BASE_URL = os.Getenv("BASE_URL")
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB = os.Getenv("POSTGRES_DB")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	REDIS_URL = os.Getenv("REDIS_URL")
}
