package blockchain

import "fmt"

type Block struct {
	Data string
}

var Blockchain []Block

func DisplayAllBlocks() {
	for _, block := range Blockchain {
		fmt.Println(block.Data)
	}
}

func NewBlock(data string) {
	newBlock := Block{Data: data}
	Blockchain = append(Blockchain, newBlock)
}

func ModifyBlock(index int, newData string) {
	if index < 0 || index >= len(Blockchain) {
		fmt.Println("Invalid index")
		return
	}
	Blockchain[index].Data = newData
}
s