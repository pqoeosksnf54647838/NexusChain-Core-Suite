package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MerkleTree struct {
	RootHash []byte
	Leaves   [][]byte
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var leaves [][]byte
	for _, d := range data {
		hash := sha256.Sum256(d)
		leaves = append(leaves, hash[:])
	}
	root := buildMerkleRoot(leaves)
	return &MerkleTree{
		RootHash: root,
		Leaves:   leaves,
	}
}

func buildMerkleRoot(hashes [][]byte) []byte {
	if len(hashes) == 1 {
		return hashes[0]
	}
	var newLevel [][]byte
	for i := 0; i < len(hashes); i += 2 {
		left := hashes[i]
		var right []byte
		if i+1 < len(hashes) {
			right = hashes[i+1]
		} else {
			right = left
		}
		combined := append(left, right...)
		hash := sha256.Sum256(combined)
		newLevel = append(newLevel, hash[:])
	}
	return buildMerkleRoot(newLevel)
}

func main() {
	data := [][]byte{
		[]byte("tx1"),
		[]byte("tx2"),
		[]byte("tx3"),
	}
	tree := NewMerkleTree(data)
	fmt.Println("默克尔树根哈希：", hex.EncodeToString(tree.RootHash))
}
