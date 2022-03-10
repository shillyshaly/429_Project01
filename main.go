package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	handleError(err)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		handleError(err)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	req, _ := parseRequest(conn)
	sendResponse(conn, req)
}

type request struct {
	method string
	//body     []byte
	uri      string
	protocol string
}

func parseRequest(conn net.Conn) (*request, error) {
	b, _, err := bufio.NewReader(conn).ReadLine()
	req := new(request)
	handleError(err)

	sp := strings.Split(string(b), " ")
	req.method, req.uri, req.protocol = sp[0], sp[1], sp[2]

	fmt.Println("Method: " + req.method)
	fmt.Println("URI: " + req.uri)
	fmt.Println("Protocol: " + req.protocol)

	return req, nil
}

func sendResponse(conn net.Conn, req *request) {
	content, _ := os.ReadFile("www" + req.uri)
	dne, err := os.ReadFile("www/404.html")
	handleError(err)

	if _, err := os.Stat("www" + req.uri); os.IsNotExist(err) {
		b := []byte("HTTP/1.1 404 DoesNotExit\r\n\r\n" + string(dne))
		conn.Write(b)
	} else {
		b := []byte("HTTP/1.1 200 OK\r\n\r\n" + string(content))
		conn.Write(b)
	}
	conn.Close()

}

func handleError(err error) {
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		os.Exit(2)
	}
}

//accept()  done
//parseRequest()  done
//generateResponse(___)
//sendResponse()
//close()
