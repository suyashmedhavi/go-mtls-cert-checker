package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
)

func main() {
	serverMode := flag.Bool("server", false, "-server")
	clientMode := flag.Bool("client", false, "-client")
	certFile := flag.String("certFile", "", "-certFile /path/to/file")
	keyFile := flag.String("keyFile", "", "-keyFile /path/to/file")
	trustFile := flag.String("trustFile", "", "-certFile /path/to/file (optional, uses system trust otherwise)")
	serverName := flag.String("serverName", "", "-serverName name (required with -client)")
	port := flag.Int("port", 8080, "-port number (optional, uses 8080)")
	flag.Parse()

	if (!*serverMode && !*clientMode) || (*serverMode && *clientMode) {
		colorPrintf(red, "Exactly one of -server or -client required\n\n")
		setColor(red)
		flag.PrintDefaults()
		resetColor()
		return
	}

	if len(*certFile) == 0 || len(*keyFile) == 0 {
		colorPrintf(red, "Certificate or key file path not given\n\n")
		setColor(red)
		flag.PrintDefaults()
		resetColor()
		return
	}

	if *clientMode {
		if len(*serverName) == 0 {
			colorPrintf(red, "-serverName required with -client\n\n")
			setColor(red)
			flag.PrintDefaults()
			resetColor()
			return
		}
	}

	certPool, caPool, err := initCerts(*certFile, *keyFile, *trustFile)
	if err != nil {
		return
	}

	if *serverMode {
		err = startServer(*port, certPool, caPool)
		if err != nil {
			colorPrintf(red, "%s", err)
			return
		}
	} else {
		err = startClient(*port, certPool, caPool, *serverName)
		if err != nil {
			colorPrintf(red, "%s", err)
			return
		}
	}
}

func initCerts(certFile, keyFile, trustFile string) (certPool tls.Certificate, caPool *x509.CertPool, err error) {
	colorPrintf(white, "Reading certs...\n")
	certPool, err = readKeyPair(certFile, keyFile)
	if err != nil {
		colorPrintf(red, "%s", err)
		return
	}
	if len(trustFile) == 0 {
		colorPrintf(white, "Using system trust...\n")
		caPool, err = x509.SystemCertPool()
		if err != nil {
			colorPrintf(red, "%s", err)
			return
		}
	} else {
		colorPrintf(white, "Reading trust...\n")
		caPool, err = readTrust(trustFile)
		if err != nil {
			colorPrintf(red, "%s", err)
			return
		}
	}
	return
}
