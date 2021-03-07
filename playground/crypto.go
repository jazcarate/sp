// Playground with crypto
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey

	fmt.Println("Public key: ", base64.StdEncoding.EncodeToString(publicKey.N.Bytes()))

	message := []byte("message to be signed")
	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Signature: %x\n", signature)

	err = rsa.VerifyPKCS1v15(&publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return
	}

	fmt.Println("signature verified")
}
