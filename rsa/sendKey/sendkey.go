package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
)

       var salt = []byte("gocows")

func main() {

	respByte, err := ioutil.ReadFile("../key")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	block, _ := pem.Decode([]byte(respByte))

	rsaPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	ciphertext, err := rsa.EncryptOAEP(sha256.New(),
		rand.Reader,
		rsaPublicKey.(*rsa.PublicKey),
		[]byte("I like cows more than gophers"), //message
		salt)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	//testDecryption(fmt.Sprintf("%s", ciphertext))

        var buf bytes.Buffer
	buf.Write(ciphertext)
	_, err = http.Post("http://127.0.0.1:8080/", "raw", &buf)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Print("Message Sent")

	return

}

func testDecryption(ciphertext string) {
	respByte, err := ioutil.ReadFile("../pkey")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	blocked, _ := pem.Decode([]byte(respByte))

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(blocked.Bytes)
	if err != nil {
		fmt.Printf("%s", err)
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(),
		rand.Reader,
		rsaPrivateKey,
		[]byte(ciphertext),
		salt)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Printf("%s", plaintext)

}
