package main

import (
	"fmt"
	"net/http"
	"time"
)

func race(site, site_two string) string {

	start_site := time.Now()
	http.Get(site)
	finish_site := time.Since(start_site)

	start_site_two := time.Now()
	http.Get(site_two)
	finish_site_two := time.Since(start_site_two)

	if finish_site_two > finish_site {
		return site
	}
	return site_two

}

func main() {

	result := race("www.google.com", "www.facebook.com")
	fmt.Println(result)

}
