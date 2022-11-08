package main

import (
	"log"
	"os"

	"bitbucket.org/jojocoders/microkit/server"
	"bitbucket.org/jojocoders/microkit/tracer"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	godotenv.Load()
	router := getRoutes()
	url := os.Getenv("JAEGER_ENDPOINT")

	tracer.InitOtel(url, "Open Tabungan Emas Service", "v1.0.0", "development")

	server.Serve("8001", router)
}
