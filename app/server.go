package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	out := "HTTP/1.1 200 OK\r\n\r\n"

	in := make([]byte, 4096)
	content, err := connection.Read(in)
	if err != nil {
		fmt.Println("error accepting connection")
		os.Exit(1)
	}

	// nc -v 127.0.0.1 4221
	// opens the channel to send anything (string) from your terminal to
	// this golang server!
	/**
	1. start this golang server
	2. run nc -v 127.0.0.1 4221
	3. type anything in your nc terminal
	4. golang will receive it!
	*/
	fmt.Println(string(in))
	fmt.Println(content)

	connection.Write([]byte(out))
}
