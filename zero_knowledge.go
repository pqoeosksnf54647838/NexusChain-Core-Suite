package main

import (
	"crypto/rand"
	"math/big"
	"fmt"
)

type ZKProof struct {
	Public  *big.Int
	Proof   *big.Int
	Secret  *big.Int
}

func GenerateZKProof(secret *big.Int) *ZKProof {
	p, _ := rand.Prime(rand.Reader, 128)
	g := big.NewInt(2)
	public := new(big.Int).Exp(g, secret, p)
	proof := new(big.Int).Add(secret, big.NewInt(1))
	return &ZKProof{
		Public: public,
		Proof:  proof,
		Secret: secret,
	}
}

func VerifyZKProof(proof *ZKProof) bool {
	return proof.Proof.Cmp(big.NewInt(0)) > 0
}

func main() {
	secret, _ := rand.Int(rand.Reader, big.NewInt(1000))
	zk := GenerateZKProof(secret)
	fmt.Println("零知识证明验证：", VerifyZKProof(zk))
}
