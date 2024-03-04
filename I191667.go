package main

import (
	"fmt"
)

// Block represents a simple data structure for a block
type Block struct {
	ID   int
	Data string
}

// Blockchain represents a collection of blocks
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block with given data
func NewBlock(id int, data string) *Block {
	return &Block{
		ID:   id,
		Data: data,
	}
}

// ModifyBlock modifies the data of a block with given ID
func (bc *Blockchain) ModifyBlock(id int, newData string) error {
	for _, block := range bc.Blocks {
		if block.ID == id {
			block.Data = newData
			return nil
		}
	}
	return fmt.Errorf("block with ID %d not found", id)
}

// DisplayAllBlocks prints all blocks in the blockchain
func (bc *Blockchain) DisplayAllBlocks() {
	fmt.Println("Blocks in the blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Block %d: %s\n", block.ID, block.Data)
	}
}

func main() {
	// Create a new blockchain
	blockchain := &Blockchain{}

	// Add some blocks to the blockchain
	blockchain.Blocks = append(blockchain.Blocks, NewBlock(1, "Data of Block 1"))
	blockchain.Blocks = append(blockchain.Blocks, NewBlock(2, "Data of Block 2"))
	blockchain.Blocks = append(blockchain.Blocks, NewBlock(3, "Data of Block 3"))

	// Display all blocks in the blockchain
	blockchain.DisplayAllBlocks()

	// Modify data of Block 2
	err := blockchain.ModifyBlock(2, "Modified data of Block 2")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Display all blocks in the blockchain again
	blockchain.DisplayAllBlocks()
}
