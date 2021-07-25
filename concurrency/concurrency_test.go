package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := checkwebsite(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {

		t.Errorf("got %v, want %v", got, want)
	}

}

func mockWebsiteChecker(url string) bool {
	time.Sleep(2 * time.Second)
	return url != "waat://furhurterwe.geds"
}
