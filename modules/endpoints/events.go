package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pterm/pterm"
)

func EventsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, "Correct method was used.")
		
	} else {
		wrongMethodError := Error{
			ErrorCode:    http.StatusMethodNotAllowed,
			ErrorMessage: "GET is the only accepted Method for this endpoint.",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)

		wrongMethodErrorJson, err := json.Marshal(wrongMethodError)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		} else {
			_, err := w.Write(wrongMethodErrorJson)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			}
		}
	}
}
