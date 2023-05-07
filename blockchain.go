package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{} //it is a map whose keys are strings and values are any type.
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) calculateHash() string {
	// Convert data to json
	data, _ := json.Marshal(b.data)
	// Create block that is including previous hash + data + timestamp + pow
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	// Return hash as String
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	fmt.Println("minsing hash is:" + b.hash)
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

// The first block does not have any value and previous hash and data becasue it is the first block in the network.
func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now()}
	return Blockchain{
		genesisBlock: genesisBlock,
		chain:        []Block{genesisBlock},
		difficulty:   difficulty}
}

// Create transaction and new bloc and mine the new block based on the previous block and finally add it to the chains.
func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	// Get the last block from the chain to get previous hash.
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	// Add new block to the chain
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currenctBlock := b.chain[i+1]
		// Calculate hash for every block and compared them and also check the previous block hash.
		if currenctBlock.hash != currenctBlock.calculateHash() || currenctBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {

	// Create new blockchain and set difficulty to be 3.
	blockchain := CreateBlockchain(3)
	fmt.Println(blockchain)

	// Add new transactions
	blockchain.addBlock("MohammadHossein", "Ali", 1000)
	blockchain.addBlock("John", "lawyer", 750)

	// Check blockchain is valid or not.
	// fmt.Println(blockchain.isValid())

	fmt.Println(blockchain.chain)

}
