package main

import (
	"context"
	"crypto/rand"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// https://github.com/libp2p/go-libp2p/blob/master/examples/pubsub/chat/main.go

func main() {
	r := rand.Reader
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)

	if err != nil {
		panic(err)
	}

	node, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"), libp2p.Identity(prvKey))

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	ps, err := pubsub.NewGossipSub(ctx, node)

	if err != nil {
		panic(err)
	}

	if err := SetupDiscovery(node); err != nil {
		panic(err)
	}

	ui := NewBlockChainUI(ps)

	ui.Run()

	if err := node.Close(); err != nil {
		panic(err)
	}
}
