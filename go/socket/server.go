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

func SocketServer() {
	listen, err := net.Listen(CONN_TYPE, net.JoinHostPort(CONN_HOST, CONN_PORT))
	if err != nil {
		log.Fatalf("Socket listen port %s failed: %v", CONN_PORT, err)
	}
	defer listen.Close()

	log.Printf("Socket listen port %s", CONN_PORT)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept failed: %s", err)
			continue
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	defer conn.Close()

	ip, _, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("new client: %s", ip)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("Error Reading: %s", err)
	}

	log.Printf("received: %s", string(buf[:n]))

	conn.Write([]byte("Message received."))

}

func main() {
	SocketServer()
}
