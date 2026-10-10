package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type memoryTransport struct{}

func (memoryTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Header: make(http.Header), Body: io.NopCloser(strings.NewReader("remote ok")), Request: r}, nil
}
func main() {
	client := &http.Client{Transport: memoryTransport{}, Timeout: time.Second}
	resp, err := client.Get("https://example.test/data")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1024))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode, string(body))
}
