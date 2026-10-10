package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func slow(w http.ResponseWriter, r *http.Request) {
	select {
	case <-r.Context().Done():
		http.Error(w, "canceled", 408)
	default:
		w.WriteHeader(204)
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req := httptest.NewRequest("GET", "/", nil).WithContext(ctx)
	rec := httptest.NewRecorder()
	slow(rec, req)
	fmt.Println(rec.Code)
}
