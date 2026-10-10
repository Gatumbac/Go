package main

import (
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
)

type broken struct{}

func (broken) Color(context.Context) (string, error) { return "", errors.New("private detail") }
func TestColor(t *testing.T) {
	for _, tc := range []struct {
		name string
		dep  Lookup
		want int
		body string
	}{{"success", fixed{}, 200, "blue"}, {"failure", broken{}, 503, "unavailable\n"}} {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			colorHandler(tc.dep).ServeHTTP(rec, httptest.NewRequest("GET", "/color", nil))
			if rec.Code != tc.want || rec.Body.String() != tc.body {
				t.Fatalf("got %d %q", rec.Code, rec.Body.String())
			}
			if strings.Contains(rec.Body.String(), "private") {
				t.Fatal("leaked detail")
			}
		})
	}
}
