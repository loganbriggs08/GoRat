package endpoints

import (
	"fmt"
	"net/http"
)

func EventNew(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint reached.")
}
