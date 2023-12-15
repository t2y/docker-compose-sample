package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from go server")
	})
	slog.Info("start http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
