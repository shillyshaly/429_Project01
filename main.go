package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

//request struct to save items from the first line of request
type request struct {
	method   string
	uri      string
	protocol string
}

func main() {
	//listen on port 8080
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	handleError(err)
	defer listen.Close()

	//keep connection open
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

//parse request and return request struct
func parseRequest(conn net.Conn) (*request, error) {
	//read from connection
	b, _, err := bufio.NewReader(conn).ReadLine()

	//create a new request struct
	req := new(request)
	handleError(err)

	//split first line
	sp := strings.Split(string(b), " ")
	req.method, req.uri, req.protocol = sp[0], sp[1], sp[2]

	//printing first line of request
	fmt.Println("Method: " + req.method)
	fmt.Println("URI: " + req.uri)
	fmt.Println("Protocol: " + req.protocol)

	return req, nil
}

//take request and form response
func sendResponse(conn net.Conn, req *request) {
	//save content from user chosen file if exists
	content, _ := ioutil.ReadFile("www" + req.uri)

	//save 404 file
	dne, err := ioutil.ReadFile("www/404.html")
	handleError(err)

	//for GET or HEAD request
	if req.method == "GET" || req.method == "HEAD" {

		//if file doesn't exist
		if _, err := os.Stat("www" + req.uri); os.IsNotExist(err) {
			//save 404 page info
			info, _ := os.Stat("www/404.html")
			len := strconv.FormatInt(info.Size(), 10)

			//send response header
			b := []byte("HTTP/1.1 404 DoesNotExit\r\nContent-Length: " + len + "\r\nServer: cihttp\r\n\r\n" + string(dne))
			conn.Write(b)
		} else {
			//save found page info
			info, _ := os.Stat("www" + req.uri)
			fmt.Println(info.ModTime())
			len := strconv.FormatInt(info.Size(), 10)

			//send response header
			b := []byte("HTTP/1.1 200 OK\r\nContent-Length:" + len + "\r\nServer: cihttp\r\n\r\n" + string(content))
			conn.Write(b)
		}
	}
	conn.Close()
}

func handleError(err error) {
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		os.Exit(2)
	}
}

//accept()
//parseRequest()
//generateResponse(___)
//sendResponse()
//close()
