package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Esperado status 200, mas obteve %v", status)
	}

	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("Esperado '%v', mas obteve '%v'", expected, rr.Body.String())
	}
}
