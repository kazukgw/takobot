package main

import (
	"io"
	"net/http"
	"os"
)

func ServeStats() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "takobot live!!")
	})
	http.ListenAndServe(":"+port, nil)
}
