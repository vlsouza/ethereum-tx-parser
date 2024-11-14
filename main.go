package main

import (
	"ethereum-tx-parser/api"
	"ethereum-tx-parser/internal/parser"
)

func main() {
	// Inicializa o parser Ethereum
	ethParser := parser.NewEthereumParser()

	// Inicializa o servidor HTTP
	server := api.NewServer(ethParser)

	// Inicia o servidor na porta 8080
	server.StartHTTPServer("8080")
}
