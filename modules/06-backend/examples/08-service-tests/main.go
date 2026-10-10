package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Lookup interface {
	Color(context.Context) (string, error)
}
type fixed struct{}

func (fixed) Color(context.Context) (string, error) { return "blue", nil }
func colorHandler(dep Lookup) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, err := dep.Color(r.Context())
		if err != nil {
			http.Error(w, "unavailable", 503)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, name)
	})
}
func main() {
	rec := httptest.NewRecorder()
	colorHandler(fixed{}).ServeHTTP(rec, httptest.NewRequest("GET", "/color", nil))
	fmt.Println(rec.Code, rec.Body.String())
}
