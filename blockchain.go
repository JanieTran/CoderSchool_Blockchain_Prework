package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	fmt.Println("******TODO: IMPLEMENT InitBlockchain!******")
	spew.Dump(Blockchain)
	// Fill me in, noble warrior.

	genesis := Block{
		"data",
		time.Now().Unix(),
		[]byte{},
		[]byte{},
	}

	genesis.Hash = genesis.calculateHash()

	Blockchain = append(Blockchain, genesis)
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	fmt.Println("******TODO: IMPLEMENT NewBlock!******")
	newBlock := Block{
		Data:      data,
		Timestamp: time.Now().Unix(),
		PrevHash:  oldBlock.Hash,
		Hash:      []byte{},
	}

	newBlock.Hash = newBlock.calculateHash()

	return newBlock
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	fmt.Println("******TODO: IMPLEMENT AddBlock!******")
	spew.Dump(Blockchain)
	// Fill me in, brave wizard.
	if b.verifyHash(Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, b)
	} else {
		return errors.New("Invalid block")
	}

	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}

func (b *Block) verifyHash(prevBlock Block) bool {
	prevHash := prevBlock.calculateHash()
	return bytes.Equal(prevHash, b.PrevHash)
}
