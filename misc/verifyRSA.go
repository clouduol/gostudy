// VerifyRSA verify a file and its signature using a certificate
package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	certPath string = "/home/captain/workspace/dockerImageAuth/pki/certs/system/system.crt"
	filePath string = "/home/captain/workspace/gostudy/misc/file.txt"
	sigPath  string = "/home/captain/workspace/gostudy/misc/signature"
)

func main() {
	// read cert file
	fmt.Println("read cert file")
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}
	// format pem cert to der bytes
	fmt.Println("parse pem format cert to der bytes")
	block, _ := pem.Decode(certBytes)
	if block == nil || block.Type != "CERTIFICATE" {
		fmt.Printf("err from func main: failed to parse pem to der\n")
		os.Exit(1)
	}
	//fmt.Printf("cert block:\n%v\n", block)
	// parse cert
	fmt.Println("parse cert")
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}

	// get public key
	fmt.Println("get public key")
	pubkey := cert.PublicKey.(*rsa.PublicKey)
	fmt.Printf("public key:\n%v\n", pubkey)

	// read file contents
	fmt.Println("read file contents")
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}
	// compute sha256 hash value of the file contets, return type is [32]byte, not a slice
	fmt.Println("compute sha256 hash value of the file contents")
	hashed := sha256.Sum256(fileBytes)

	// read signature contents
	sigBytes, err := ioutil.ReadFile(sigPath)
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}

	// verify the signature
	fmt.Println("verify the signature")
	err = rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hashed[:], sigBytes)
	if err != nil {
		fmt.Printf("signature verification: failed\n")
		os.Exit(1)
	}
	fmt.Printf("signature verification: OK\n")
}
