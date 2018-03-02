package p2p

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/superliuwr/go-naive-chain/lib/service"
)

// Agent defines a P2P agent
type Agent struct {
	port              string
	blockchainService *service.BlockchainService
}

// NewAgent returns a new Agent instance
func NewAgent(port string, blockchainService *service.BlockchainService) *Agent {
	return &Agent{
		port:              port,
		blockchainService: blockchainService,
	}
}

// Start starts a P2P agent
func (a *Agent) Start() error {
	log.Println("Listening for P2P connections on ", a.port)

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go a.handleConn(conn)
	}
}

func (a *Agent) handleConn(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "Enter a new payload:")

	scanner := bufio.NewScanner(conn)

	go func() {
		for {
			time.Sleep(30 * time.Second)

			output, err := json.Marshal(a.blockchainService.Chain())
			if err != nil {
				log.Fatal(err)
			}

			io.WriteString(conn, "\n"+string(output))
			io.WriteString(conn, "\nEnter a new payload:")
		}
	}()

	for scanner.Scan() {
		payload, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("%v not a number: %v", scanner.Text(), err)
			continue
		}

		_, err = a.blockchainService.Add(payload)
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nEnter a new payload:")
	}
}
