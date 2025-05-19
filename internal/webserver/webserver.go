package webserver

import (
	"fmt"
	"net/http"
	"os"
)

func Start() error {
	port := os.Getenv("WEB_SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	handlers()

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}

	return err
}

func handlers() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
}
