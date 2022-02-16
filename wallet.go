package main

import "os"

const WalletFile = "wallet.txt"

type Wallet struct {
	Key string
}

func NewWallet(s string) *Wallet {
	return &Wallet{
		Key: s,
	}
}

func NewWalletFromFile() (*Wallet, error) {
	data, err := os.ReadFile(WalletFile)
	if err != nil {
		return nil, err
	}
	return &Wallet{
		Key: string(data),
	}, nil
}

func (w *Wallet) SaveToFile() error {
	return os.WriteFile(WalletFile, []byte(w.Key), 0644)
}
