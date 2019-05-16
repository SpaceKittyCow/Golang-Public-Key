package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"time"
	"fmt"
	"log"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	label := []byte("secret")
	secret := "I like cows more than gophers"

	log.Printf("Encypting Secret")
	log.Printf("%s",secret)

	ciphertext, err := rsa.EncryptOAEP(sha256.New(),
		rand.Reader,
		&privateKey.PublicKey,
		[]byte(secret), //message
		label)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	time.Sleep(5)
	log.Printf("Ciphertext")
	log.Printf("%s", []byte(ciphertext))
	plaintext, err := rsa.DecryptOAEP(sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		label)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	log.Printf("Plaintext")
	log.Printf("%s", plaintext)

	return
}
