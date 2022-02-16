package main

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const (
	TopicDefault                = "blockchain-room"
	BlockchainMessageBufferSize = 256
)

type BlockChainRoom struct {
	Ctx      context.Context
	Ps       *pubsub.PubSub
	Topic    *pubsub.Topic
	Sub      *pubsub.Subscription
	Self     peer.ID
	Wallet   *Wallet
	Messages chan *BlockchainMessage
}

type BlockchainMessage struct {
	Message      string
	SenderID     string
	SenderWallet string
}

func JoinBlockChainRoom(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, w *Wallet) (*BlockChainRoom, error) {
	topic, err := ps.Join(TopicDefault)
	if err != nil {
		return nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	br := &BlockChainRoom{
		Ctx:      ctx,
		Ps:       ps,
		Topic:    topic,
		Sub:      sub,
		Self:     selfID,
		Wallet:   w,
		Messages: make(chan *BlockchainMessage, BlockchainMessageBufferSize),
	}

	go br.readLoop()

	return br, nil
}

func (br *BlockChainRoom) Publish(message string) error {
	m := BlockchainMessage{
		Message:      message,
		SenderID:     br.Self.Pretty(),
		SenderWallet: br.Wallet.Key,
	}
	msgBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return br.Topic.Publish(br.Ctx, msgBytes)
}

func (br *BlockChainRoom) readLoop() {
	for {
		msg, err := br.Sub.Next(br.Ctx)
		if err != nil {
			return
		}
		if msg.ReceivedFrom == br.Self {
			continue
		}
		bm := new(BlockchainMessage)
		err = json.Unmarshal(msg.Data, bm)
		if err != nil {
			continue
		}
		br.Messages <- bm
	}
}
