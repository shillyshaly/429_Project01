package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleError(err error) {
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		os.Exit(2)
	}
}

func testShit() {
	ln, err := net.Listen("tcp", ":8080")
	handleError(err)

	for {
		conn, err := ln.Accept()
		handleError(err)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		msg, _ := bufio.NewReader(conn).ReadString(' ')
		msg = strings.TrimSuffix(msg, "\n")
		fmt.Println("fuck: " + msg)
	}

}

func main() {
	testShit()
}
