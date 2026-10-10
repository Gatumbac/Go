package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, "Hi "+r.PathValue("name"))
	})
	return mux
}
func main() {
	rec := httptest.NewRecorder()
	handler().ServeHTTP(rec, httptest.NewRequest("GET", "/greet/Uma", nil))
	fmt.Println(rec.Code, rec.Body.String())
}
