package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const infuraURL = "https://sepolia.infura.io/v3/1282f4141a8c4631820a048de222c249"

// const ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
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
	fmt.Println("Block number ", block.Number())
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress("0x34198aaD335426bd66d40842F4A13096d58d3B56"), nil)
	if err != nil {
		log.Fatalf("Error to get a balance: %v", err)
	}
	fmt.Println("Balance ", balance)
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println("fBalance ", fBalance)
	value := new(big.Float).Quo(fBalance, big.NewFloat((math.Pow10(18))))
	fmt.Println("value ", value)
}
