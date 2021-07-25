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
	return url != "waat://furhurterwe.geds"
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := []string{"a url"}
	// for i := 0; i < len(urls); i++ {

	// 	urls = append(urls, "a url")

	// }

	for i := 0; i < b.N; i++ {
		checkwebsite(slowWebsiteChecker, urls)
	}

}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
