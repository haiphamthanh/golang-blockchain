package blockchain

/* DEFINATIONs */
// Make a BlockChain struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	Blocks []*Block
}

/* BODY OF CODE */
// Make deriveHash
// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

// Create a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	buildData(block)
	return block
}

// Add block to chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Create an inititial block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Build first block in chain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

/* Data construction functions */
func buildData(block *Block) {
	// Oldway: block.DeriveHash()

	// Use proof
	pow := NewProof(block)
	nonce, hash := pow.findNonceAndHash()

	block.Hash = hash
	block.Nonce = nonce
}
