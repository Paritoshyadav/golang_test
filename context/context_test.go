package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseHeader struct {
	written bool
}

func (s *SpyResponseHeader) Header() http.Header {
	s.written = true
	return nil

}

func (s *SpyResponseHeader) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")

}

func (s *SpyResponseHeader) WriteHeader(statusCode int) {
	s.written = true

}

type Teststore struct {
	response string
	t        *testing.T
}

func (store *Teststore) Fetch(ctx context.Context) (string, error) {

	data := make(chan string, 1)

	go func() {
		var result string
		for _, v := range store.response {
			select {
			case <-ctx.Done():
				store.t.Log("got cancelled in between")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(v)

			}

		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()

	case res := <-data:
		return res, nil
	}

}

// func (store *Teststore) assertNoCancel() {
// 	store.t.Helper()
// 	if store.cancelled {
// 		store.t.Error("store was not told to cancel")
// 	}

// }

// func (store *Teststore) assertCancel() {
// 	store.t.Helper()
// 	if !store.cancelled {
// 		store.t.Error("store was told to cancel but did not")
// 	}

// }

func TestContext(t *testing.T) {
	want := "testing"

	t.Run("Getting correct response", func(t *testing.T) {

		val := &Teststore{response: want, t: t}
		s := Server(val)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		respone := httptest.NewRecorder()

		s.ServeHTTP(respone, request)

		if respone.Body.String() != want {
			t.Errorf("want %s but get %s", want, respone.Body.String())
		}

	})

	t.Run("Cancel context", func(t *testing.T) {

		val := &Teststore{response: want, t: t}
		s := Server(val)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		requestctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(requestctx)

		respone := &SpyResponseHeader{}

		s.ServeHTTP(respone, request)

		if respone.written {
			t.Errorf("should not be written")
		}

	})

}
