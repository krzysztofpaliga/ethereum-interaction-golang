package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// const infuraURL = "https://sepolia.infura.io/v3/1282f4141a8c4631820a048de222c249"
const ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), ganacheURL)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	} else {
		log.Print("Ether client created successfully")
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}
	fmt.Println(block.Number())
}
