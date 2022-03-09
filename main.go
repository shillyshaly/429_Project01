package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	dir, _ := generateResponse(req)
	sendResponse(dir, conn)
}

type request struct {
	method string // GET, POST, etc.
	//header textproto.MIMEHeader
	body     []byte
	uri      string // The raw URI from the request
	protocol string // "HTTP/1.1"
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

func generateResponse(req *request) (string, error) {
	files, err := ioutil.ReadDir("www")
	handleError(err)

	var dir string

	for _, file := range files {
		if req.method == "GET" || req.method == "HEAD" {
			if strings.Contains("/"+file.Name(), req.uri) {
				dir = file.Name()
			}
		}
	}
	return dir, nil
}

func sendResponse(dir string, conn net.Conn) {
	content, err := ioutil.ReadFile("www/index.html")
	handleError(err)
	conn.Write(content)
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
