package main

import (
	"context"
	"io/ioutil"
	"log"

	todo "ethereum-interaction-golang/11-solidity/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	const todoContractAdd = "0x40c93b1CF8cB1ed67D9029739E9cdFdAc898Cc70"

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

	file, err := ioutil.ReadFile("./05-keystore/UTC--2024-10-25T16-12-49.129947329Z--cd086624760f8f77dce849730eda1cf2d88c239f")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	key, err := keystore.DecryptKey(file, "123qwe")
	if err != nil {
		log.Fatalf("Error decrypting key: %v", err)
	}
	log.Println("Key ", hexutil.Encode(crypto.FromECDSA(key.PrivateKey)))

	todo, err := todo.NewTodo(common.HexToAddress(todoContractAdd), client)
	if err != nil {
		log.Fatalf("Error binding to todo: %v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), key.Address)
	if err != nil {
		log.Fatalf("Error getting nonce: %v", err)
	}
	log.Println("PendingNonce ", nonce)
	suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error getting gas price suggestion: %v", err)
	}
	log.Println("Suggested gas price ", suggestedGasPrice)
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error getting chain id: %v", err)
	}
	log.Println("Chain id ", chainId)

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		log.Fatalf("Error creating Transactor: %v", err)
	}
	auth.GasPrice = suggestedGasPrice
	auth.GasLimit = uint64(3_000_000)
	// auth.Nonce = big.NewInt(int64(nonce))

	tx, err := todo.Add(auth, "Call smart contract method")
	if err != nil {
		log.Fatalf("Error calling smart contract: %v", err)
	}
	log.Println("Issued smart contract tx ", tx.Hash())

	callOpts := &bind.CallOpts{
		From:    key.Address,
		Pending: false,
		Context: context.Background(),
	}
	todos, err := todo.List(callOpts)
	if err != nil {
		log.Fatalf("Error getting todos: %v", err)
	}
	log.Printf("todos %v\n", todos)

}
