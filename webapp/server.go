package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(rw, "20")

}
