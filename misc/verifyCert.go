// VerifyCert verify a cert using root ca
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	certPath string = "/home/captain/workspace/dockerImageAuth/pki/certs/system/system.crt"
	//caPath   string = "/home/captain/pki/ca/root-ca.crt"
	caPath string = "/home/captain/workspace/dockerImageAuth/pki/ca/root-ca.crt"
)

func parseCert(certPath string) (*x509.Certificate, error) {
	// read cert file
	fmt.Printf("cert path:%s\n", certPath)
	fmt.Println("read cert file")
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("err from func parseCert: %v\n", err)
	}
	// format pem cert to der bytes
	fmt.Println("parse pem format cert to der bytes")
	block, _ := pem.Decode(certBytes)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("err from func parseCert: failed to parse pem to der\n")
	}
	//fmt.Printf("cert block:\n%v\n", block)
	// parse cert
	fmt.Println("parse cert")
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("err from func parseCert: %v\n", err)
	}
	return cert, nil
}

func main() {
	// parse cert and ca
	fmt.Println("parse cert and ca")
	cert, err := parseCert(certPath)
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}
	ca, err := parseCert(caPath)
	if err != nil {
		fmt.Printf("err from func main: %v\n", err)
		os.Exit(1)
	}

	// create new cert pool
	fmt.Println("create new cert pool")
	roots := x509.NewCertPool()
	roots.AddCert(ca)

	// create VerifyOptions
	fmt.Println("create new VerifyOptions")
	opts := x509.VerifyOptions{
		Roots: roots,
	}

	// verify cert
	fmt.Println("verify cert")
	if _, err := cert.Verify(opts); err != nil {
		fmt.Printf("err from func main: %v\n", err)
		fmt.Println("Cert Verification: fail")
		os.Exit(1)
	}
	fmt.Println("Cert Verification: OK")
}
