package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK\n"))
	})
	if err := http.ListenAndServe(":18080", nil); err != nil {
		panic(err)
	}

}
