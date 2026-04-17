package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type ValidateBlock struct {
	Block
	Difficulty int
}

func (vb *ValidateBlock) VerifyHash() bool {
	record := fmt.Sprintf("%d%s%s%s%d%d", vb.Index, vb.Timestamp, vb.Data, vb.PrevHash, vb.Nonce, vb.Difficulty)
	hash := sha256.Sum256([]byte(record))
	calculated := hex.EncodeToString(hash[:])
	return calculated == vb.Hash
}

func (vb *ValidateBlock) VerifyDifficulty() bool {
	prefix := ""
	for i := 0; i < vb.Difficulty; i++ {
		prefix += "0"
	}
	return vb.Hash[:vb.Difficulty] == prefix
}

func (vb *ValidateBlock) VerifyPrevious(prevBlock Block) bool {
	return vb.PrevHash == prevBlock.Hash && vb.Index == prevBlock.Index+1
}

func (vb *ValidateBlock) FullValidate(prevBlock Block) bool {
	return vb.VerifyHash() && vb.VerifyDifficulty() && vb.VerifyPrevious(prevBlock)
}
