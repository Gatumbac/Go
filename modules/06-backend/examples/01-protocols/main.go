package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	raw := "GET /health HTTP/1.1\r\nHost: example.test\r\n\r\n"
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(raw)))
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	fmt.Println(req.Method, req.URL.Path, req.Host)
}
