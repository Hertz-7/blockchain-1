package main

import (
	"./blockchain"
	"fmt"
)

func main() {
	blockchain.NewBlock("First Block")
	blockchain.NewBlock("Second Block")
	blockchain.DisplayAllBlocks()

	blockchain.ModifyBlock(1, "Modified Second Block")
	fmt.Println("After modification:")
	blockchain.DisplayAllBlocks()
}
