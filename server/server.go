package main

import (
	"net/http"

	"github.com/NotKatsu/Otter/endpoints"
)

func main() {
	http.HandleFunc("/connection/new", endpoints.ConnectionNew())
	
}
