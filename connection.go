package main

import (
	"io"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/multiformats/go-multiaddr"
)

func MakeHost(randomess io.Reader) (host.Host, error) {
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, randomess)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	sourceMultiAdrr, _ := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/0")

	return libp2p.New(
		libp2p.ListenAddrs(sourceMultiAdrr),
		libp2p.Identity(prvKey),
	)
}
