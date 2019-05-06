// crypto/rand.Reader is a good source of entropy for blinding the RSA
// operation.
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	salt := []byte("gocows")

	ciphertext, err := rsa.EncryptOAEP(sha256.New(),
		rand.Reader,
		&privateKey.PublicKey,
		[]byte("I like cows more than gophers"), //message
		salt)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		salt)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Printf("%s", plaintext)

	return
}

    //use 
    /*privateKeyFile, err := os.Create(privateKeyPath)
    defer privateKeyFile.Close()
    if err != nil {
        return err
    }*/
    /*privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
    if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
        return err
    }*/
   /* pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
    if err != nil {
        return err
    }
*/
