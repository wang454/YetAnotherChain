package YetAnotherChain

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

func calculateHash(blk Block) string{
	blockData := string(blk.Index) + string(blk.Timestamp) + blk.PrevBlockHash + blk.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func generateNewBlock(prevBlk Block, data string) Block {
	newBlk := Block{}
	newBlk.Index = prevBlk.Index + 1
	newBlk.Timestamp = time.Now().Unix()
	newBlk.PrevBlockHash = prevBlk.CurrBlockHash
	newBlk.Data = data
	newBlk.CurrBlockHash = calculateHash(newBlk)
	return newBlk
}

func generateGenesisBlock() Block {
	preBlk := Block{}
	preBlk.Index = -1
	preBlk.CurrBlockHash = ""
	return generateNewBlock(preBlk, "Genesis Block")
}