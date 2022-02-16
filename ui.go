package main

import (
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/rivo/tview"
)

// https://github.com/libp2p/go-libp2p/blob/master/examples/pubsub/chat/ui.go

type BlockChainUI struct {
	Room      *BlockChainRoom
	app       *tview.Application
	mainPanel *tview.Flex
	mnList    *tview.List
	ps        *pubsub.PubSub
	Wallet    *Wallet
	doneCh    chan struct{}
}

func NewBlockChainUI(br *BlockChainRoom, ps *pubsub.PubSub, wallet *Wallet) *BlockChainUI {
	app := tview.NewApplication()

	mainPanel := tview.NewFlex()
	mainPanel.SetBorder(true).SetTitle(fmt.Sprintf("Blockchain - %s", wallet.Key))

	doneCh := make(chan struct{}, 1)

	ui := &BlockChainUI{
		Room:      br,
		app:       app,
		mainPanel: mainPanel,
		ps:        ps,
		doneCh:    doneCh,
		Wallet:    wallet,
	}

	ui.build()
	return ui
}

func (ui *BlockChainUI) build() {
	logBox := tview.NewTextView()
	logBox.SetDynamicColors(true)
	logBox.SetBorder(true)
	logBox.SetTitle(fmt.Sprintf("Mensagens%s", ""))
	logBox.SetChangedFunc(func() {
		ui.app.Draw()
	})

	peersList := tview.NewTextView()
	peersList.SetBorder(true)
	peersList.Box.SetTitle("Conectados")
	peersList.SetChangedFunc(func() {
		ui.app.Draw()
	})

	mnList := tview.NewList().
		AddItem("Transferir", "Transferência de valores entre carteiras", 't', nil).
		AddItem("Sair", "Sair do Sistema", 's', func() {
			ui.app.Stop()
			ui.doneCh <- struct{}{}
		})
	mnList.SetBorder(true)
	mnList.SetTitle("Menu")

	ui.mnList = mnList

	panel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ui.mainPanel, 30, 1, false).
		AddItem(logBox, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(mnList, 40, 1, true).
		AddItem(panel, 0, 1, false).
		AddItem(peersList, 40, 1, false)

	ui.app.SetRoot(flex, true).EnableMouse(true)
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
