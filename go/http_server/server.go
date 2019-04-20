package main

import (
	"io"
	"net/http"
)

func main() {

	// static
	fs := http.FileServer(http.Dir("/tmp/"))
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", fs))

	// dynamic
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "welcome to my website!")
	}
	http.HandleFunc("/ping", helloHandler)

	if err := http.ListenAndServe("192.168.150.235:80", nil); err != nil {
		panic(err)
	}
}
