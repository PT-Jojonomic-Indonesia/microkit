package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PT-Jojonomic-Indonesia/microkit/database/postgres"
	"github.com/PT-Jojonomic-Indonesia/microkit/server"
	"github.com/PT-Jojonomic-Indonesia/microkit/service/example-postgres/entity"
	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s dbname=%s port=%s user=%s password=%s TimeZone=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"),
	)

	config := gorm.Config{}

	log.Printf("\n[info] connecting to postgres with dsn : %s", dsn)
	if err := postgres.Init(dsn, &config); err != nil {
		log.Panic(err)
	}
	log.Println("\n[info] connected to postgres")

	log.Printf("\n[info] running auto migrate")
	if err := postgres.Migrate(entity.Nasabah{}); err != nil {
		log.Panic(err)
	}
	log.Println("\n[info] successfull migrate all table")

	router := getRoutes()
	url := os.Getenv("JAEGER_ENDPOINT")

	tracer.InitOtel(url, "Example Postgres Service", "v1.0.0", "development")

	server.Serve("8001", router)
}
