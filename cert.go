package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func readKeyPair(certPath, keyPath string) (certs tls.Certificate, err error) {
	certs, err = tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return
	}
	for i, cert := range certs.Certificate {
		x509Cert, err := x509.ParseCertificate(cert)
		if err != nil {
			return certs, err
		}
		boxPrint(green, []string{
			fmt.Sprintf("Certificate %d", i+1),
			fmt.Sprintf("CN: %s", x509Cert.Subject.CommonName),
			fmt.Sprintf("SANs: %v", x509Cert.DNSNames),
			// fmt.Sprintf("Extensions: %v", x509Cert.Extensions),
			fmt.Sprintf("Issuer: %v", x509Cert.Issuer.CommonName),
		})
	}
	return
}

func readTrust(trustFile string) (caPool *x509.CertPool, err error) {
	caCertData, err := ioutil.ReadFile(trustFile)
	if err != nil {
		err = fmt.Errorf("trust: readFile: %w", err)
		return
	}
	caPool, ok := getTrustPool(caCertData)
	if !ok {
		err = fmt.Errorf("trust: parse cert: none found")
		return
	}
	return caPool, nil
}

func getTrustPool(pemCerts []byte) (caPool *x509.CertPool, ok bool) {
	i := 0
	caPool = x509.NewCertPool()
	for len(pemCerts) > 0 {
		var block *pem.Block
		block, pemCerts = pem.Decode(pemCerts)
		if block == nil {
			break
		}
		if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
			continue
		}
		certBytes := block.Bytes
		cert, err := x509.ParseCertificate(certBytes)
		if err != nil {
			continue
		}
		caPool.AddCert(cert)
		boxPrint(green, []string{
			fmt.Sprintf("Trust %d", i+1),
			fmt.Sprintf("CN: %s", cert.Subject.CommonName),
			// fmt.Sprintf("Extensions: %v", x509Cert.Extensions),
			fmt.Sprintf("Issuer: %v", cert.Issuer.CommonName),
		})
		ok = true
		i++
	}
	return
}
