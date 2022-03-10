The program runs with 
    go run main.go

After the server is running, going to localhost:8080/index.html brings up the index page.
Going to localhost:8080 seems to not return anything but an empty page.  

Once the connection is made and listening, the connection is read, parsed and split into
the request method, uri and protocol.  Then a response is generated depending on whether
the page exists.  If the page doesn't exist, the 404 page is returned.  Otherwise, the requested
page is returned.