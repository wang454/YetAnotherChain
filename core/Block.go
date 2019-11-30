package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct{
	Index int64 // block index
	Timestamp int64 // block timestamp
	PrevBlockHash string // previous block hash
	CurrBlockHash string // current block hash

	Data string // block data
}

func CalculateHash(blk Block) string{
	blockData := string(blk.Index) + string(blk.Timestamp) + blk.PrevBlockHash + blk.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = preBlock.CurrBlockHash
	newBlock.Data = data
	newBlock.CurrBlockHash = CalculateHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.CurrBlockHash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}