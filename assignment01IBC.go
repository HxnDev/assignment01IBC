package assignment01IBC

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
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
	tran := inputBlock.Data.Transactions
	tranintostring := strings.Join(tran[:], ",")

	hash := sha256.Sum256([]byte(tranintostring))
	hashh := hex.EncodeToString(hash[:])

	return hashh
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
	i := 1
	for newHead != nil {
		fmt.Println("Block Number = ", i)
		fmt.Print("Transactions = ")
		fmt.Println(newHead.Data.Transactions)
		fmt.Print("Hash = ")
		fmt.Println(newHead.CurrentHash)
		fmt.Print("Hash of Block ", i-1, " = ")
		fmt.Printf(newHead.PrevHash)
		i++
		fmt.Print("\n\n")
		newHead = newHead.PrevPointer
	}
}

// VerifyChain verifies the blockchain for illegal transaction
func VerifyChain(chainHead *Block) {
	for c := chainHead; c != nil; c = c.PrevPointer {
		hashc := CalculateHash(c)
		if c.PrevPointer != nil {
			hashp := CalculateHash(c.PrevPointer)
			if hashp != c.PrevHash || hashc != c.CurrentHash {
				fmt.Println("Blockchain is compromised")
				return
			}
		}
		if hashc != c.CurrentHash {
			fmt.Println("Blockchain is compromised")
			return
		}
	}
	fmt.Println("Blockchain Verified")
	return
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
		current = current.PrevPointer
	}
}
