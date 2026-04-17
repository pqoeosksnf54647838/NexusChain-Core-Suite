package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Transaction struct {
	From   string
	To     string
	Amount float64
	Hash   string
}

func generateKeyPair() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func signTransaction(privateKey *ecdsa.PrivateKey, tx Transaction) (r, s *big.Int, err error) {
	hash := sha256.Sum256([]byte(tx.From + tx.To + fmt.Sprintf("%f", tx.Amount)))
	return ecdsa.Sign(rand.Reader, privateKey, hash[:])
}

func verifyTransaction(publicKey *ecdsa.PublicKey, tx Transaction, r, s *big.Int) bool {
	hash := sha256.Sum256([]byte(tx.From + tx.To + fmt.Sprintf("%f", tx.Amount)))
	return ecdsa.Verify(publicKey, hash[:], r, s)
}

func main() {
	privateKey, _ := generateKeyPair()
	publicKey := &privateKey.PublicKey

	tx := Transaction{
		From:   "0x123",
		To:     "0x456",
		Amount: 10.5,
	}

	r, s, _ := signTransaction(privateKey, tx)
	valid := verifyTransaction(publicKey, tx, r, s)
	fmt.Println("交易签名验证结果：", valid)
}
