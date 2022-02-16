package main

import (
	"flag"
)

// https://github.com/libp2p/go-libp2p/blob/master/examples/pubsub/chat/main.go

func main() {
	walletKey := flag.String("wallet", "", "Wallet key")

	flag.Parse()

	var (
		wallet *Wallet
		err    error
	)

	// usuário não passou a chave da carteira
	if *walletKey == "" {
		// tenta ler do arquivo
		wallet, err = NewWalletFromFile()

		if err != nil {
			*walletKey = RandString(8)
			wallet = NewWallet(*walletKey)
			wallet.SaveToFile()
		}

	}

	// r := rand.Reader
	// prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)

	// if err != nil {
	// 	panic(err)
	// }

	// node, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"), libp2p.Identity(prvKey))

	// if err != nil {
	// 	panic(err)
	// }

	// ctx := context.Background()

	// ps, err := pubsub.NewGossipSub(ctx, node)

	// if err != nil {
	// 	panic(err)
	// }

	// if err := SetupDiscovery(node); err != nil {
	// 	panic(err)
	// }

	// br, err := JoinBlockChainRoom(ctx, ps, node.ID(), wallet)
	// if err != nil {
	// 	panic(err)
	// }

	// ui := NewBlockChainUI(br, ps, wallet)

	// ui.Run()

	// if err := node.Close(); err != nil {
	// 	panic(err)
	// }

	bc := NewBlockchain()
	bg := NewBlockGenesis()
	bc.AddBlock(bg)
	b := NewBlock([]byte(bg.Hash()))
	b.AddTransaction(NewTransaction(TransactionAdd, wallet, 1))
	bc.AddBlock(b)
	bc.SaveToFile()
}
