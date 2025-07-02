# 🔐 EVM Wallet Generator

[![Go Reference](https://pkg.go.dev/badge/github.com/exorich-lab/EVM-GEN/v1.svg)](https://pkg.go.dev/github.com/exorich-lab/EVM-GEN/v1)
[![GitHub](https://img.shields.io/badge/GitHub-Repository-blue?style=for-the-badge&logo=github)](https://github.com/exorich-lab/EVM-GEN)
[![Go Report Card](https://goreportcard.com/badge/github.com/exorich-lab/EVM-GEN)](https://goreportcard.com/report/github.com/exorich-lab/EVM-GEN)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[![GitHub](https://img.shields.io/badge/GitHub-Repository-blue?style=for-the-badge&logo=github)](https://github.com/exorich-lab/EVM-GEN)


A simple and secure tool for generating Ethereum-compatible (EVM) wallets using BIP39 mnemonics and BIP44 derivation paths. The tool can be used both as a Go module in your applications or as a standalone CLI tool.

## ✨ Features

- Generate secure Ethereum wallets with 24-word mnemonic phrases
- Derive Ethereum addresses using BIP44 standard derivation path (m/44'/60'/0'/0/0)
- Export wallet information to encrypted files
- Use as a Go module in your projects or as a standalone CLI tool
- Cross-platform compatibility (Windows, macOS, Linux)
- Open source and MIT licensed

## 📋 Prerequisites

- Go 1.16 or later (only required if building from source or using as a module)
- Git (for cloning the repository)
- For CLI usage, Go 1.16+ is only needed for installation

## 🚀 Quick Start

### Option 1: Use as a Go Module

```bash
go get github.com/exorich-lab/EVM-GEN/v1@latest
```

Example usage in your Go code:

```go
package main

import (
	"fmt"
	"log"

	"github.com/exorich-lab/EVM-GEN/v1/pkg/evmwallet"
)

func main() {
	// Generate a new wallet
	wallet, err := evmwallet.GenerateWallet()
	if err != nil {
		log.Fatal(err)
	}

	// Print wallet information
	fmt.Println("Address:", wallet.Address)
	fmt.Println("Mnemonic:", wallet.Mnemonic)
	
	// Save to file
	filename, err := wallet.SaveToFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wallet saved to:", filename)
}
```

### Option 2: Install as a CLI Tool

```bash
# Install the latest version
go install github.com/exorich-lab/EVM-GEN/v1/cmd/evmwallet@latest

# Run the tool
evmwallet
```

### Option 3: Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/exorich-lab/EVM-GEN.git
   cd EVM-GEN
   ```

2. Build and install:
   ```bash
   go install ./cmd/evmwallet
   ```

3. Run the tool:
   ```bash
   evmwallet
   ```

## 💻 CLI Usage

Generate a new wallet:

```bash
evmwallet
```

This will:
1. Generate a new 24-word mnemonic
2. Derive an Ethereum address and private key
3. Save the details to a timestamped file in the `wallets` directory
4. Display the information in the console

### 📝 Example Output

```
Mnemonic: word1 word2 ... word24

Ethereum Address: 0x...
Private Key: 0x...

Wallet information has been saved to wallet_20230702_123456.txt
```

## 🔒 Security Notes

- **Never share your mnemonic phrase or private key with anyone**
- Always store your mnemonic phrase in a secure location
- The generated wallet information is saved to a file in the current directory - make sure to secure or delete this file appropriately
- For production use, consider using hardware wallets or other more secure key management solutions

## 📜 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

Contributions are welcome! 🙌 Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## ⚠️ Disclaimer

This software is provided "as is" without warranty of any kind. Use at your own risk. The developers are not responsible for any loss of funds or other damages caused by using this software.
