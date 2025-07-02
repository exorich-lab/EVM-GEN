package evmwallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// Wallet represents an Ethereum wallet with its mnemonic, private key, and address
type Wallet struct {
	Mnemonic   string
	PrivateKey *ecdsa.PrivateKey
	Address    string
}

// GenerateWallet creates a new Ethereum wallet
func GenerateWallet() (*Wallet, error) {
	// Generate a random mnemonic with 256 bits of entropy (24 words)
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, fmt.Errorf("failed to generate entropy: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, fmt.Errorf("failed to generate mnemonic: %v", err)
	}

	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "") // No passphrase for simplicity

	// Derive the master key from the seed
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to derive master key: %v", err)
	}

	// Derive path m/44'/60'/0'/0/0 - standard Ethereum derivation path
	purpose, err := masterKey.NewChildKey(0x8000002C) // 44' - BIP44 purpose
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose: %v", err)
	}

	coinType, err := purpose.NewChildKey(0x8000003C) // 60' - Ethereum coin type
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type: %v", err)
	}

	account, err := coinType.NewChildKey(0x80000000) // 0' - Account 0
	if err != nil {
		return nil, fmt.Errorf("failed to derive account: %v", err)
	}

	change, err := account.NewChildKey(0x80000000) // 0 - External chain
	if err != nil {
		return nil, fmt.Errorf("failed to derive change: %v", err)
	}

	addressIndex, err := change.NewChildKey(0x80000000) // 0 - Address index
	if err != nil {
		return nil, fmt.Errorf("failed to derive address index: %v", err)
	}

	privateKey, err := crypto.ToECDSA(addressIndex.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA private key: %v", err)
	}

	// Get the public key and derive the Ethereum address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &Wallet{
		Mnemonic:   mnemonic,
		PrivateKey: privateKey,
		Address:    address,
	}, nil
}

// SaveToFile saves the wallet information to a file
func (w *Wallet) SaveToFile() (string, error) {
	dir := "wallets"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create wallets directory: %v", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join(dir, fmt.Sprintf("wallet_%s.txt", timestamp))

	privateKeyBytes := crypto.FromECDSA(w.PrivateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	content := fmt.Sprintf(`Mnemonic: %s

Ethereum Address: %s
Private Key: 0x%s

IMPORTANT: Keep this file safe and never share your private key or mnemonic with anyone!`,
		w.Mnemonic, w.Address, privateKeyHex)

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write wallet to file: %v", err)
	}

	return filename, nil
}
