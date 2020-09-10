package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)


type BlockChain struct {
	Blocks []*Block
}

// block type
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// for creating a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// for adding a block; taken prev block & then append a new block
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// decide which block contain valid data
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}