package main

import (
	"log"
	"time"
	
	"ethereum-tx-parser/api"
	"ethereum-tx-parser/internal/parser"
	"ethereum-tx-parser/internal/rpcclient"
)

func main() {

	// Initialize ethereum parser
	ethParser := parser.NewEthereumParser(
		rpcclient.New("https://ethereum-rpc.publicnode.com"), // can be an ENV var
	)

	// Starts a go routine to update the current block address
	go func() {
		for {
			err := ethParser.UpdateCurrentBlock()
			if err != nil {
				log.Printf("Erro ao atualizar o bloco atual: %v", err)
			}
			time.Sleep(10 * time.Second) // the time window can also be an ENV var
		}
	}()

	// Initialize http server
	server := api.NewServer(ethParser)

	// Initializer server at port 8080
	server.StartHTTPServer("8080")
}
