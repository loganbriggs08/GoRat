package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/pterm/pterm"
)

type Error struct {
	ErrorCode    uint64 `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func ConnectionNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ComputerMACAddress := r.Header.Get("ComputerMACAddress")
		ComputerOS := r.Header.Get("ComputerOS")
		ComputerName := r.Header.Get("ComputerName")

		if ComputerMACAddress == "" || ComputerName == "" || ComputerOS == "" {

		}

	} else {
		NewMethodError := Error{
			ErrorCode:    http.StatusMethodNotAllowed,
			ErrorMessage: "POST is the only accepted Method for this endpoint.",
		}

		w.WriteHeader(http.StatusMethodNotAllowed)

		NewReturnMethodError, err := json.Marshal(NewMethodError)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		} else {
			_, err := w.Write(NewReturnMethodError)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			}
		}
	}
}
