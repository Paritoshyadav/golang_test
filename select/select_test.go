package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		site1 := getslowurl(5 * time.Second)

		site2 := getslowurl(4 * time.Second)
		defer site1.Close()
		defer site2.Close()

		slowUrl := site1.URL
		fastUrl := site2.URL

		got, err := race(slowUrl, fastUrl)
		if err != nil {
			t.Fatalf("Error should not be occur %s", err)
		}

		want := fastUrl

		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		site1 := getslowurl(8 * time.Millisecond)

		site2 := getslowurl(10 * time.Millisecond)
		defer site1.Close()
		defer site2.Close()

		slowUrl := site1.URL
		fastUrl := site2.URL

		_, err := configrace(slowUrl, fastUrl, 5*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})

}

func getslowurl(timer time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(timer)

		w.WriteHeader(http.StatusOK)

	}))
}
