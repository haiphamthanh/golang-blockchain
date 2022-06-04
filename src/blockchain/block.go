package blockchain

import (
	"bytes"
	"crypto/sha256"
)

/* DEFINATIONs */
// TODO: 1. Define Block struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// TODO: 4. Make a BlockChain struct
type BlockChain struct {
	blocks []*Block
}

/* BODY OF CODE */
// TODO: 2. Make deriveHash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// TODO: 3. Create a block
func CreateBlock(data string, prevHash []byte) *Block {
	// 1. Create a new block with empty hash byte
	block := &Block{[]byte{}, []byte(data), prevHash}
	// 2. Fill your hash
	block.DeriveHash()
	// 3. Return the block
	return block
}

// TODO: 5. Add block to chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// TODO: 6. Make a inititial block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// TODO: 7. Make Init BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
