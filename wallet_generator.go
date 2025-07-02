package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Generate a random mnemonic with 256 bits of entropy (24 words)
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		log.Fatalf("Failed to generate entropy: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatalf("Failed to generate mnemonic: %v", err)
	}

	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "") // No passphrase for simplicity

	// Derive the master key from the seed
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatalf("Failed to derive master key: %v", err)
	}

	// Derive path m/44'/60'/0'/0/0 - standard Ethereum derivation path
	// m/44'/60'/0'/0/0 is the derivation path for the first account in MetaMask
	// 44' - BIP44 purpose
	// 60' - Ethereum coin type
	// 0' - Account 0
	// 0 - External chain (for change addresses, 0 for normal addresses, 1 for change)
	// 0 - Address index (0 for first address)

	// Derive m/44'
	purpose, err := masterKey.NewChildKey(0x8000002C) // 44' - BIP44 purpose
	if err != nil {
		log.Fatalf("Failed to derive purpose: %v", err)
	}

	// Derive m/44'/60'
	coinType, err := purpose.NewChildKey(0x8000003C) // 60' - Ethereum coin type
	if err != nil {
		log.Fatalf("Failed to derive coin type: %v", err)
	}

	// Derive m/44'/60'/0'
	account, err := coinType.NewChildKey(0x80000000) // 0' - Account 0
	if err != nil {
		log.Fatalf("Failed to derive account: %v", err)
	}

	// Derive m/44'/60'/0'/0
	externalChain, err := account.NewChildKey(0) // 0 - External chain
	if err != nil {
		log.Fatalf("Failed to derive external chain: %v", err)
	}

	// Derive m/44'/60'/0'/0/0
	addressKey, err := externalChain.NewChildKey(0) // 0 - Address index
	if err != nil {
		log.Fatalf("Failed to derive address key: %v", err)
	}

	// Convert to ECDSA private key
	privateKey, err := crypto.ToECDSA(addressKey.Key)
	if err != nil {
		log.Fatalf("Failed to convert to ECDSA: %v", err)
	}

	// Get the public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// Prepare wallet info for saving
	var sb strings.Builder
	sb.WriteString("=== EVM Wallet ===\n\n")
	
	// Split mnemonic into words and number them
	words := strings.Fields(mnemonic)
	for i, word := range words {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, word))
	}

	sb.WriteString("\nPublic Address: " + address + "\n")
	sb.WriteString("\nPrivate Key: " + hex.EncodeToString(crypto.FromECDSA(privateKey)) + "\n")
	sb.WriteString("\nCreated: " + time.Now().Format("2006-01-02 15:04:05"))

	// Create wallets directory if it doesn't exist
	err = os.MkdirAll("wallets", 0755)
	if err != nil {
		log.Fatalf("Failed to create wallets directory: %v", err)
	}

	// Save to file
	filename := filepath.Join("wallets", "wallet_"+time.Now().Format("20060102_150405")+".txt")
	err = os.WriteFile(filename, []byte(sb.String()), 0600)
	if err != nil {
		log.Fatalf("Failed to save wallet: %v", err)
	}

	fmt.Printf("Wallet successfully created and saved to %s\n", filename)
	fmt.Println("\nIMPORTANT: Keep this file safe and never share your private key or mnemonic with anyone!")
}
