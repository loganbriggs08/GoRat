package endpoints

import (
	"fmt"
	"net/http"
)

func EventsGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
