package main

import (
	"fmt"
	"log"
	"os"

	"bitbucket.org/jojocoders/microkit/database/db2"
	"bitbucket.org/jojocoders/microkit/server"
	"bitbucket.org/jojocoders/microkit/tracer"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	godotenv.Load()

	dsn := fmt.Sprintf(
		"HOSTNAME=%s;DATABASE=%s;PORT=%s;UID=%s;PWD=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
	)
	log.Printf("\n[info] connecting to DB2 with dsn : %s", dsn)
	if err := db2.Init(dsn); err != nil {
		log.Panic(err)
	}
	defer db2.DB.Close()
	log.Println("\n[info] connected to DB2")

	router := getRoutes()
	url := os.Getenv("JAEGER_ENDPOINT")

	tracer.InitOtel(url, "Example DB2 Service", "v1.0.0", "development")

	server.Serve("8001", router)
}
