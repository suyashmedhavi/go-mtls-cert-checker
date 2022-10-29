package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
)

func startServer(port int, certs tls.Certificate, trust *x509.CertPool) error {
	address := fmt.Sprintf(":%d", port)
	config := tls.Config{
		Certificates: []tls.Certificate{certs},
		// ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          trust,
		MinVersion:         tls.VersionTLS13,
		InsecureSkipVerify: true,
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	listener = tls.NewListener(listener, &config)
	colorPrintf(yellow, "Server listening on %s\n", listener.Addr())
	for {
		colorPrintf(yellow, "Waiting for connection...\n")
		conn, err := listener.Accept()
		if err != nil {
			colorPrintf(red, "%s\n", err)
			continue
		}
		err = conn.(*tls.Conn).Handshake()
		if err != nil {
			colorPrintf(red, "%s\n", err)
			continue
		}
		go func(conn1 net.Conn) {
			for {
				data := make([]byte, 999999)
				n, err := conn1.Read(data)
				if err != nil {
					colorPrintf(red, "%s\n", err)
					return
				}
				colorPrintf(yellow, "Read: %s\n", string(data[:n]))
			}
		}(conn)
	}
}
