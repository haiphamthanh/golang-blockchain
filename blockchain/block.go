package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

/* DEFINATIONs */
// TODO: 4. Make a BlockChain struct
type BlockChain struct {
	Blocks []*Block
}

// TODO: 1. Define Block struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

/* BODY OF CODE */
// TODO: 2. Make deriveHash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// TODO: 3. Create a block
// func CreateBlock(data string, prevHash []byte) *Block {
// 	// 1. Create a new block with empty hash byte
// 	block := &Block{[]byte{}, []byte(data), prevHash}
// 	// 2. Fill your hash
// 	block.DeriveHash()
// 	// 3. Return the block
// 	return block
// }
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

// TODO: 5. Add block to chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// TODO: 6. Make a inititial block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// TODO: 7. Make Init BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// Other
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash = *intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
