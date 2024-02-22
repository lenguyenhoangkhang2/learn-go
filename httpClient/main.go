package main

import "net/http"

const serverPort = 3333

func main() {
	go func() {
		mux := http.NewServeMux()
	}()
}
