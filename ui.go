package main

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/rivo/tview"
)

// https://github.com/libp2p/go-libp2p/blob/master/examples/pubsub/chat/ui.go

type BlockChainUI struct {
	app    *tview.Application
	ps     *pubsub.PubSub
	doneCh chan struct{}
}

func NewBlockChainUI(ps *pubsub.PubSub) *BlockChainUI {
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("Blockchain - UFCoin")
	app.SetRoot(box, true)

	return &BlockChainUI{
		app: app,
		ps:  ps,
	}
}

func (ui *BlockChainUI) Run() error {
	go ui.handleEvents()
	defer ui.End()
	return ui.app.Run()
}

func (ui *BlockChainUI) handleEvents() {
	for {
		select {
		case <-ui.doneCh:
			return
		}
	}
}

func (ui *BlockChainUI) End() {
	ui.doneCh <- struct{}{}
}
