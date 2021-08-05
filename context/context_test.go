package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Teststore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (store *Teststore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return store.response

}

func (store *Teststore) Cancel() {

	store.cancelled = true

}

func (store *Teststore) assertNoCancel() {
	store.t.Helper()
	if store.cancelled {
		store.t.Error("store was not told to cancel")
	}

}

func (store *Teststore) assertCancel() {
	store.t.Helper()
	if !store.cancelled {
		store.t.Error("store was told to cancel but did not")
	}

}

func TestContext(t *testing.T) {
	want := "testing"

	t.Run("Getting correct response", func(t *testing.T) {

		val := &Teststore{response: want, cancelled: false, t: t}
		s := Server(val)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		respone := httptest.NewRecorder()

		s.ServeHTTP(respone, request)

		if respone.Body.String() != want {
			t.Errorf("want %s but get %s", want, respone.Body.String())
		}
		val.assertNoCancel()

	})

	t.Run("Cancel context", func(t *testing.T) {

		val := &Teststore{response: want, cancelled: false, t: t}
		s := Server(val)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		requestctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(requestctx)

		respone := httptest.NewRecorder()

		s.ServeHTTP(respone, request)

		val.assertCancel()

	})

}
