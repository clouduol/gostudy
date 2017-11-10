// SignRSA demonstrates sign a file using sha256withRSAEncryption
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	prikeyPath string = "/home/captain/workspace/dockerImageAuth/pki/certs/platform/platform.key"
	filePath   string = "/home/captain/workspace/gostudy/misc/file.txt"
	sigPath    string = "/home/captain/workspace/gostudy/misc/signature"
)

func main() {
	// read private key file
	fmt.Println("read private key file")
	prikeyBytes, err := ioutil.ReadFile(prikeyPath)
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}
	// format pem private key to der bytes
	fmt.Println("parse pem format private key to der bytes")
	block, _ := pem.Decode(prikeyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Printf("err from func main: failed to parse pem to der\n")
		os.Exit(1)
	}
	encBool := x509.IsEncryptedPEMBlock(block)
	if encBool {
		fmt.Println("encryption pem block")
	} else {
		fmt.Println("non encryption pem block")
	}
	//fmt.Printf("private key block:\n%v\n", block)
	// parse private key
	fmt.Println("parse private key ")
	prikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}

	// get pseudorandom number generator
	fmt.Println("get pseudorandom number generator")
	rng := rand.Reader

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

	// sign the hash value
	fmt.Println("sign the hash value")
	signature, err := rsa.SignPKCS1v15(rng, prikey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("signature:\n %x\n", signature)

	// write signature to file
	fmt.Println("write signature to file")
	var modePerm os.FileMode = 0644
	if err := ioutil.WriteFile(sigPath, signature, modePerm); err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}

}
