package main

import (
	"log"
	"os"

	"bitbucket.org/jojocoders/microkit/server"
	"bitbucket.org/jojocoders/microkit/tracer"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	router := getRoutes()
	url := os.Getenv("JAEGER_ENDPOINT")

	tracer.InitOtel(url, "Tabungan Emas Repository", "v1.0.0", "development")

	server.Serve("8002", router)
}
