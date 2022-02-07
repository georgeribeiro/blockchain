package main

import "github.com/rivo/tview"

// https://github.com/libp2p/go-libp2p/blob/master/examples/pubsub/chat/ui.go

type BlockChainUI struct {
	app    *tview.Application
	doneCh chan struct{}
}

func NewBlockChainUI() *BlockChainUI {
	app := tview.NewApplication()

	return &BlockChainUI{
		app: app,
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
