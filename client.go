package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
)

func startClient(port int, certs tls.Certificate, trust *x509.CertPool, serverName string) error {
	address := fmt.Sprintf(":%d", port)
	config := &tls.Config{
		Certificates: []tls.Certificate{certs},
		RootCAs:      trust,
		MinVersion:   tls.VersionTLS13,
		ServerName:   serverName,
	}
	raddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return err
	}
	c := tls.Client(conn, config)
	err = c.Handshake()
	if err != nil {
		return err
	}
	defer c.Close()
	for {
		msg := uuid.New().String()
		colorPrintf(yellow, "Sending: %s\n", msg)
		_, err = c.Write([]byte(msg))
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
