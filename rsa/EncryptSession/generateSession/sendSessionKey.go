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

       var label = []byte("sessionKey")

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

	hash := sha256.New()
	hash.Size()
	//OEAP Session Key cannot be > keyByteSize - (hashSize *2) - 2
	sessionKey := make([]byte, (rsaPublicKey.(*rsa.PublicKey).Size()-(hash.Size()*2)-3))
	//populates Session with random data
	rand.Read(sessionKey)

	fmt.Printf("%v \n", []rune(string(sessionKey)))
	ciphertext, err := rsa.EncryptOAEP(sha256.New(),
		rand.Reader,
		rsaPublicKey.(*rsa.PublicKey),
		sessionKey,
		label)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	//confirmCorrectDecryption(fmt.Sprintf("%s", ciphertext))

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

func confirmCorrectDecryption(ciphertext string) {
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
		label)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Printf("%s", plaintext)

}
