package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/superliuwr/go-naive-chain/lib/service"
	"github.com/superliuwr/go-naive-chain/lib/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(run())
}

// web server
func run() error {
	blockchainService := service.NewBlockchainService()
	webServer := web.NewServer(os.Getenv("ADDR"), blockchainService)

	return webServer.Start()
}
