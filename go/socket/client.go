package main

import (
	"log"
	"net"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func SocketClient() {
	conn, err := net.Dial(CONN_TYPE, net.JoinHostPort(CONN_HOST, CONN_PORT))
	if err != nil {
		log.Fatalf("connect %s %s failed: %s", CONN_HOST, CONN_PORT, err)
	}
	defer conn.Close()

	conn.Write([]byte("lalalala"))

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])
}

func main() {
	SocketClient()
}
