package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {

	random := rand.Reader
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	caCert := generateCATemplate()
	//can place anytype of private key in 
	derBytes, err := x509.CreateCertificate(random, caCert,
		caCert, &privateKey.PublicKey, privateKey)
        if err != nil {
		fmt.Printf("%s", err)
		return
	}

       certCerFile, err := os.Create("ca.cer")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
       certCerFile.Write(derBytes)

       certPEMFile, err := os.Create("ca.pem")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
       pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
       certPEMFile.Close()

}

func generateCATemplate()(*x509.Certificate){
	//most of this function came from https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/security/x509_certificates.html
	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000) // one year
	return &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "Hannah",
			Organization: []string{"SpaceKittyCow"},
		},
		NotBefore: now,
		NotAfter:  then,

		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,

		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"github.com", "localhost"},
	}
}
