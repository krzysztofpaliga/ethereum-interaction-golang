package main

import (
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// keyStore := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	// account, err := keyStore.NewAccount("123qwe")
	// if err != nil {
	// 	log.Fatalf("Error creating account: %v", err)
	// }
	// log.Println("Account created ", account)
	file, err := ioutil.ReadFile("./keystore/UTC--2024-10-25T15-57-19.320812599Z--4bc66192738b3dafc8f381d094ffe23197f0b990")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	key, err := keystore.DecryptKey(file, "123qwe")
	if err != nil {
		log.Fatalf("Error decrypting key: %v", err)
	}
	log.Println("Key ", hexutil.Encode(crypto.FromECDSA(key.PrivateKey)))
	log.Println("Public Key ", hexutil.Encode(crypto.FromECDSAPub(&key.PrivateKey.PublicKey)))
	log.Println("Address ", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}
