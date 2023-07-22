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
		ID := r.Header.Get("ID")
		IP := r.Header.Get("IP")
		OS := r.Header.Get("OS")
		Name := r.Header.Get("Name")

		if ID == "" || OS == "" || Name == "" || IP == "" {
			NewUnauthorizedError := Error{
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: "Content is missing from request headers.",
			}

			w.WriteHeader(http.StatusUnauthorized)

			NewReturnUnauthorizedError, err := json.Marshal(NewUnauthorizedError)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			} else {
				_, err := w.Write(NewReturnUnauthorizedError)

				if err != nil {
					pterm.Fatal.WithFatal(true).Println(err)
				}
			}
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
