package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Teststore struct {
	response string
}

func (store *Teststore) Fetch() string {

	return store.response

}

func TestContext(t *testing.T) {
	want := "testing"
	val := &Teststore{response: want}
	s := Server(val)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	respone := httptest.NewRecorder()

	s.ServeHTTP(respone, request)

	if respone.Body.String() != want {

		t.Errorf("want %s but get %s", want, respone.Body.String())

	}

}
