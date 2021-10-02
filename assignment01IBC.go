package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

// BlockData is a structure containing Transactions
type BlockData struct {
	Transactions []string //Transactions in a block
}

// Block contains all the necessary information for a chain
type Block struct {
	Data        BlockData // Stores Data (Transactions)
	PrevPointer *Block    // Previous Block
	PrevHash    string    // Hash of Previous Block
	CurrentHash string    // Hash of Current Block
}

// CalculateHash : Computes the Hash of a complete block
func CalculateHash(inputBlock *Block) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%v", *inputBlock))))
}

// InsertBlock : Inserts a block into the blockchain
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	if chainHead == nil {
		var newBlock Block
		chainHead = &newBlock
		chainHead.PrevPointer = nil
		chainHead.Data = dataToInsert
		chainHead.CurrentHash = CalculateHash(chainHead)
		chainHead.PrevHash = ""
	} else {
		var newBlock Block
		newBlock.PrevPointer = chainHead
		newBlock.PrevHash = chainHead.CurrentHash
		newBlock.Data = dataToInsert
		newBlock.CurrentHash = CalculateHash(&newBlock)
		chainHead = &newBlock
	}
	return chainHead
}

// ListBlocks displays a list of every block
func ListBlocks(chainHead *Block) {
	newHead := chainHead
	for true {
		if newHead == nil {
			break
		}
		fmt.Print("HASH: ")
		fmt.Println(newHead.CurrentHash)
		fmt.Print("TRANSACTIONS: ")
		fmt.Println(newHead.Data.Transactions)
		fmt.Print("PREV HASH: ")
		fmt.Printf(newHead.PrevHash)
		fmt.Print("\n----------------------------------------\n")
		newHead = newHead.PrevPointer
	}
}

// VerifyChain verifies the blockchain for illegal transaction
func VerifyChain(chainHead *Block) {
	if chainHead != nil {
		for chainHead.PrevPointer != nil {
			if chainHead.PrevHash != CalculateHash(chainHead.PrevPointer) {
				fmt.Println("Mismatched")
				//fmt.Println(chainHead.PrevPointer.Data + " has current hash as " + CalculateHash(chainHead.PrevPointer) + " but it should be " + chainHead.PrevHash)
			} else {
				fmt.Println("Verified")
			}
			chainHead = chainHead.PrevPointer
		}
	}

	/*
	     newHead := chainHead
	   	for true {
	   		if newHead.PrevPointer == nil {
	   			break
	   		}
	   		prevhashstored := newHead.PrevHash
	   		fmt.Print("PREV HASH STORED : ")
	   		fmt.Println(prevhashstored)
	   		newHead = newHead.PrevPointer
	   	}*/
}

// ChangeBlock -> alter a transaction
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	current := chainHead

	for current != nil {
		transactions := current.Data.Transactions

		for i := range transactions {
			if transactions[i] == oldTrans {
				transactions[i] = newTrans

				current.CurrentHash = CalculateHash(current)
				break
			}
		}
	}
}
