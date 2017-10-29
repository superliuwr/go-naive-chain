package main

import (
	"flag"
	"log"

	"github.com/superliuwr/go-naive-chain/lib/service"
	"github.com/superliuwr/go-naive-chain/lib/util"	
	"github.com/superliuwr/go-naive-chain/lib/web"
)

func main() {
	webServerPort := flag.String("port", "8080", "web server port number")
	nodeID := flag.String("id", util.PseudoUUID(), "node ID")
	flag.Parse()

	// Step 1 Create a blockchain and initialize genesis block
	blockchain := service.NewBlockchain()

	// Step 2 Setup peers connection
	// Step 3 Start P2P server
	// Step 4 Start HTTP server
	httpServer := web.NewServer(webServerPort, blockchain, *nodeID)
	err := httpServer.Start()
	if err != nil {
		log.Fatal("Unable to start Web server", err)
	}
}