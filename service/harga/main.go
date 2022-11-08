package main

import (
	"log"
	"os"

	"github.com/PT-Jojonomic-Indonesia/microkit/server"
	"github.com/PT-Jojonomic-Indonesia/microkit/tracer"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	router := getRoutes()
	url := os.Getenv("JAEGER_ENDPOINT")

	tracer.InitOtel(url, "Harga Emas Service", "v1.0.0", "development")

	server.Serve("8002", router)
}
