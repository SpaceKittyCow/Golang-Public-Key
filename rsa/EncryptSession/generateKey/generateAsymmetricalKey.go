package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"os"
)

func main() {
	label := []byte("sessionKey")
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	pub, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Printf("%s", err)
	}
	pe := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pub,
	}

	file, err := os.Create("../key")
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = pem.Encode(file, pe)
	if err != nil {
		fmt.Printf("%s", err)
	}
	//createPrivateKey(privateKey)
	fmt.Print("Public Key Set, starting Server \n")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, err := ioutil.ReadAll(r.Body)
			ciphertext := fmt.Sprintf("%s", body)

			plaintext, err := rsa.DecryptOAEP(sha256.New(),
				rand.Reader,
				privateKey,
				[]byte(ciphertext),
				label)

			if err != nil {
				fmt.Printf("%s", err)
				io.WriteString(w, "Invalid Data")
				return
			}

			fmt.Printf("%v \n", []rune(string(plaintext)))
			return
		} else {
			fmt.Print("Please send POST")
		}
	})
	http.ListenAndServe(":8080", nil)
}

func createPrivateKey(privateKey *rsa.PrivateKey){

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	file, err := os.Create("../pkey")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	err = pem.Encode(file, block)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Print("PrivateKeySet")
	return
}

