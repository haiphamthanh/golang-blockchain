package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

var Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

/* BODY OF CODE */
// Take the data in the block
func NewProof(b *Block) *ProofOfWork {
	taget := big.NewInt(1)
	taget.Lsh(taget, uint(256-Difficulty))

	pow := &ProofOfWork{b, taget}
	return pow
}

// Find valid nonce in hash algorithm
// Create a counter (nonce) which starts at 0
// https://biastek.com/vi/blockchain-co-ban-giai-thuat-dong-thuan-pow/
func (pow *ProofOfWork) findNonceAndHash() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := pow.initData(nonce)

		// Create a hash of the data plus the counter
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()
	return nonce, hash[:]
}

// Check the hash to see if it meets a set of requirements
// Requirements:
// The first few bytes must contain 0s
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.initData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash = *intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

/* Supportors*/
// Combine prev + data + nonce + difficulty
func (pow *ProofOfWork) initData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			toHex(int64(nonce)),
			toHex(int64(Difficulty)),
		},
		[]byte{})
	return data
}

func toHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
