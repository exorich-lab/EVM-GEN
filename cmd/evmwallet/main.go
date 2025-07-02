package main

import (
	"fmt"
	"log"
	"os"

	"github.com/exorich-lab/EVM-GEN/v1/pkg/evmwallet"
)

func main() {
	// Generate a new wallet
	wallet, err := evmwallet.GenerateWallet()
	if err != nil {
		log.Fatalf("Failed to generate wallet: %v", err)
	}

	// Save wallet to file
	filename, err := wallet.SaveToFile()
	if err != nil {
		log.Fatalf("Failed to save wallet: %v", err)
	}

	// Print wallet information
	privateKeyBytes := wallet.PrivateKey.D.Bytes()
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)

	fmt.Println("\n✅ Wallet successfully created!")
	fmt.Println("\n🔑 Mnemonic:", wallet.Mnemonic)
	fmt.Println("\n📬 Address:", wallet.Address)
	fmt.Println("\n🔒 Private Key: 0x" + privateKeyHex)
	fmt.Printf("\n💾 Wallet information has been saved to %s\n", filename)

	// Print security warning in red
	fmt.Println("\n\x1b[31m⚠️  IMPORTANT: Keep this file safe and never share your private key or mnemonic with anyone!\x1b[0m")

	// Keep the window open on Windows
	if _, err := os.Stat("C:\\Windows"); err == nil {
		fmt.Println("\nPress Enter to exit...")
		fmt.Scanln()
	}
}
