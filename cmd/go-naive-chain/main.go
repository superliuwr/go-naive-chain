package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/superliuwr/go-naive-chain/lib/p2p"
	"github.com/superliuwr/go-naive-chain/lib/service"
	"github.com/superliuwr/go-naive-chain/lib/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	blockchainService := service.NewBlockchainService()

	go func() {
		log.Fatal(runWebServer(blockchainService))
	}()

	log.Fatal(runP2PServer(blockchainService))
}

// web server
func runWebServer(blockchainService *service.BlockchainService) error {
	webServer := web.NewServer(os.Getenv("WEB_PORT"), blockchainService)

	return webServer.Start()
}

// p2p server
func runP2PServer(blockchainService *service.BlockchainService) error {
	p2pServer := p2p.NewAgent(os.Getenv("P2P_PORT"), blockchainService)

	return p2pServer.Start()
}
