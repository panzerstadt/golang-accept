package main

import (
	"fmt"
	"strings"

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

	in := make([]byte, 4096)
	_, err = connection.Read(in)
	if err != nil {
		fmt.Println("error accepting connection")
		os.Exit(1)
	}

	inputString := string(in)
	rows := strings.Split(inputString, "\r\n")

	// for idx, row := range rows {

	// 	fmt.Println("row number " + fmt.Sprint(idx))
	// 	fmt.Println(row)
	// }

	f := strings.Split(rows[0], " ")
	verb := f[0]
	route := f[1]
	httpVersion := f[2]

	fmt.Println(verb, route, httpVersion)

	if route == "/" {
		out := "HTTP/1.1 200 OK\r\n\r\n"
		connection.Write([]byte(out))
		return
	}

	if strings.Contains(route, "/echo/") {
		body := strings.Replace(route, "/echo/", "", 1)

		// HTTP status line
		out := "HTTP/1.1 200 OK\r\n\r\n"
		// headers
		out += "Content-Type: text/plain\r\n"
		out += fmt.Sprintf("Content-Length: %d\r\n", len(body))
		out += "\r\n" // end of headers
		// body
		out += body

		connection.Write([]byte(out))
		return
	}

	out := "HTTP/1.1 404 Not Found\r\n\r\n"
	connection.Write([]byte(out))
}
