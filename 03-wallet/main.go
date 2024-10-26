package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const infuraURL = "https://sepolia.infura.io/v3/1282f4141a8c4631820a048de222c249"
const privateKeyHex = ""

// const ganacheURL = "http://localhost:8545"

func main() {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Error instantiating private key: %v", err)
	}
	log.Println("Private key ", privateKey)

	pData := crypto.FromECDSA(privateKey)
	fmt.Println("pData ", pData)
	encoded := hexutil.Encode(pData)
	fmt.Println("encoded ", encoded)

	encodedPublicKey := hexutil.Encode(crypto.FromECDSAPub(&privateKey.PublicKey))
	fmt.Println("encoded public key ", encodedPublicKey)

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Println("address", address)
}
