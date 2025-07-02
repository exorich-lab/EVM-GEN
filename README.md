# 🔐 EVM Wallet Generator

[![GitHub](https://img.shields.io/badge/GitHub-Repository-blue?style=for-the-badge&logo=github)](https://github.com/exorich-lab/EVM-GEN)


A simple command-line tool for generating Ethereum-compatible (EVM) wallets using BIP39 mnemonics and BIP44 derivation paths. The tool generates a mnemonic phrase and derives an Ethereum address along with its private key.

## ✨ Features

- Generates a 24-word BIP39 mnemonic phrase
- Derives an Ethereum address using the standard derivation path (m/44'/60'/0'/0/0)
- Saves wallet information to a timestamped text file
- Displays wallet details in the console
- Cross-platform compatibility (Windows, macOS, Linux)

## 📋 Prerequisites

- Go 1.16 or later (only required if building from source)
- Git (for cloning the repository)

## ⚙️ Installation

### 📥 Option 1: Download Pre-built Binary

1. Download the latest release from the [Releases](https://github.com/yourusername/evm-wallet-generator/releases) page
2. For Windows, download `evm-wallet-generator-windows-amd64.exe`
3. Rename it to `evm-wallet-generator` (or `evm-wallet-generator.exe` on Windows)
4. Make it executable (on Unix-like systems):
   ```bash
   chmod +x evm-wallet-generator
   ```

### 🔧 Option 2: Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/exorich-lab/EVM-GEN.git
   cd EVM-GEN
   ```

2. Build the binary:
   ```bash
   go build -o evm-wallet-generator
   ```

## 🚀 Usage

### 💻 Basic Usage

Run the executable:

```bash
./evm-wallet-generator
```

On Windows:
```cmd
evm-wallet-generator.exe
```

### 📊 Output

The program will:
1. Generate a new 24-word mnemonic
2. Derive an Ethereum address and private key
3. Display the information in the console
4. Save the details to a file named `wallet_<timestamp>.txt` in the current directory

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
