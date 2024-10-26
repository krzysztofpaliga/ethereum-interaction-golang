package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const infuraURL = "https://sepolia.infura.io/v3/1282f4141a8c4631820a048de222c249"

func main() {
	// keyStore := keystore.NewKeyStore("./05-keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := keyStore.NewAccount("123qwe")
	// if err != nil {
	// 	log.Fatalf("Error crating keystore: %v", err)
	// }

	// _, err = keyStore.NewAccount("123qwe")
	// if err != nil {
	// 	log.Fatalf("Error crating keystore: %v", err)
	// }

	const addr1 = "cd086624760f8f77dce849730eda1cf2d88c239f"
	const addr2 = "55c5efd12561a1bfe176c393c522c59523c531f8"

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Error creating ether client: %v", err)
	}
	defer client.Close()

	balance1, err := client.BalanceAt(context.Background(), common.HexToAddress(addr1), nil)
	if err != nil {
		log.Fatalf("Error getting balance1: %v", balance1)
	}
	log.Println("balance1 ", balance1)

	balance2, err := client.BalanceAt(context.Background(), common.HexToAddress(addr2), nil)
	if err != nil {
		log.Fatalf("Error getting balance2: %v", balance2)
	}
	log.Println("balance2 ", balance2)
}
