package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

const DiscoveryServiceTag = "blockchain-ufccoin"

type DiscoverNotifee struct {
	h host.Host
}

func (d *DiscoverNotifee) HandlePeerFound(pi peer.AddrInfo) {
	err := d.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

func SetupDiscovery(h host.Host) error {
	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &DiscoverNotifee{h: h})
	return s.Start()
}
