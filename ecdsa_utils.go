package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func PublicKeyToHex(publicKey *ecdsa.PublicKey) string {
	xBytes := publicKey.X.Bytes()
	yBytes := publicKey.Y.Bytes()
	return hex.EncodeToString(append(xBytes, yBytes...))
}

func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(privateKey.D.Bytes())
}

func main() {
	priv, _ := GenerateECDSAKey()
	pub := &priv.PublicKey
	fmt.Println("私钥：", PrivateKeyToHex(priv))
	fmt.Println("公钥：", PublicKeyToHex(pub))
}
