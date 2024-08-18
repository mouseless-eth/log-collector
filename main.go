package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999" // Default port if not set
	}

	addr := "0.0.0.0:" + port
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()
	log.Println("Server is running on:", addr)

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Failed to accept conn.", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("New connection from: %s\n", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		content := scanner.Text()
		fmt.Print(content)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection %s: %v\n", conn.RemoteAddr(), err)
	}
}
