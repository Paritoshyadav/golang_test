package main

import (
	"fmt"
	"net/http"
	"time"
)

func race(site, site_two string) (string, error) {

	return configrace(site, site_two, 10*time.Second)

}

func configrace(site, site_two string, duration time.Duration) (string, error) {

	select {

	case <-ping(site):
		return site, nil

	case <-ping(site_two):
		return site_two, nil

	case <-time.After(duration):
		return "", fmt.Errorf("time out")

	}

}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {

		http.Get(url)
		close(ch)

	}()
	return ch
}

func main() {

	result, _ := race("www.google.com", "www.facebook.com")
	fmt.Println(result)

}
