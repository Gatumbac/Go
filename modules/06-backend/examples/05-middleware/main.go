package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func stamp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Header().Set("X-Course", "Go"); next.ServeHTTP(w, r) })
}
func main() {
	h := stamp(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(204) }))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	fmt.Println(rec.Code, rec.Header().Get("X-Course"))
}
