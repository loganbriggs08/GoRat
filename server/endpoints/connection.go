package endpoints

import (
	"fmt"
	"net/http"
)

func ConnectionNew(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
