package main

import "net/http"

func errPanic(_ http.ResponseWriter, _ *http.Request) error {
	panic(123)
}
