package main

import (
	"fmt"

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

	doneCh := make(chan struct{}, 1)

	logBox := tview.NewTextView()
	logBox.SetDynamicColors(true)
	logBox.SetBorder(true)
	logBox.SetTitle(fmt.Sprintf("Log: %s", "1"))
	logBox.SetChangedFunc(func() {
		app.Draw()
	})

	peersList := tview.NewTextView()
	peersList.SetBorder(true)
	peersList.Box.SetTitle("Conectados")
	peersList.SetChangedFunc(func() {
		app.Draw()
	})

	mnList := tview.NewList().
		AddItem("Transferir", "Transferência de valores entre carteiras", 't', nil).
		AddItem("Sair", "Sair do Sistema", 's', func() {
			app.Stop()
			doneCh <- struct{}{}
		})
	mnList.SetBorder(true)
	mnList.SetTitle("Menu")

	panel := tview.NewFlex().
		AddItem(mnList, 20, 1, true).
		AddItem(logBox, 0, 1, false).
		AddItem(peersList, 20, 1, false)

	app.SetRoot(panel, true).EnableMouse(true)

	return &BlockChainUI{
		app:    app,
		ps:     ps,
		doneCh: doneCh,
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
	ui.app.Stop()
}
