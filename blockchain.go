package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

type TransactionType byte

const (
	TransactionAdd    TransactionType = 'a'
	TransactionCredit TransactionType = 'c'
	TransactionDebt   TransactionType = 'd'
	BlockchainFile                    = "ledger.json"
	Target                            = 999999999
	BlockDraftSize                    = 256
)

type Transaction struct {
	ID     []byte          `json:"id"`
	Wallet []byte          `json:"wallet"`
	Type   TransactionType `json:"type"`
	Value  float64         `json:"value"`
}

type Block struct {
	Timestamp    int64          `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
	PrevHash     []byte         `json:"prevHash"`
	Draft        bool
	foundCh      chan int64
}

type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

var draftBlocks []*Block = []*Block{}

func (b *Block) Hash() string {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	transactions, _ := json.Marshal(b.Transactions)
	headers := bytes.Join([][]byte{b.PrevHash, timestamp, transactions}, []byte{})
	hash := sha256.Sum256(headers)
	return fmt.Sprintf("%x", hash)
}

func NewBlockGenesis() *Block {
	return &Block{
		Timestamp:    time.Now().Unix(),
		Transactions: []*Transaction{},
		PrevHash:     []byte{},
	}
}

func NewBlock(prevHash []byte) *Block {
	b := &Block{
		Timestamp:    time.Now().Unix(),
		Transactions: []*Transaction{},
		PrevHash:     prevHash,
		Draft:        true,
	}
	draftBlocks = append(draftBlocks, b)
	return b
}

func (b *Block) AddTransaction(t *Transaction) *Block {
	b.Transactions = append(b.Transactions, t)
	if b.Full() {
		b.Mine()
	}
	return b
}

func (b *Block) Full() bool {
	return len(b.Transactions) >= BlockDraftSize
}

func (b *Block) Mine() {
	for {
		n := RandInt()
		hasher := sha256.New()
		hasher.Write([]byte(strconv.FormatInt(n, 10)))
		h := hasher.Sum(nil)
		hi := new(big.Int).SetBytes(h)
		// encontrou o valor que é menor que o target
		if hi.Int64() < Target {
			b.Draft = false
			b.foundCh <- hi.Int64()
			return
		}
	}
}

func (bc *Blockchain) SaveToFile() error {
	data, err := json.Marshal(bc)
	if err != nil {
		return err
	}
	err = os.WriteFile(BlockchainFile, data, 0644)
	return err
}

func NewBlockchainFromFile() (*Blockchain, error) {
	data, err := os.ReadFile(BlockchainFile)
	if err != nil {
		return nil, err
	}
	var bc Blockchain
	err = json.Unmarshal(data, &bc)
	if err != nil {
		return nil, err
	}
	return &bc, err
}

func NewTransaction(typ TransactionType, w *Wallet, v float64) *Transaction {
	return &Transaction{
		ID:     []byte(RandString(32)),
		Type:   typ,
		Wallet: []byte(w.Key),
		Value:  v,
	}
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{},
	}
}

func (bc *Blockchain) AddBlock(b *Block) *Blockchain {
	bc.Blocks = append(bc.Blocks, b)
	return bc
}
