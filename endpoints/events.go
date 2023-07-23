package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/pterm/pterm"
)

func EventNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Authorization") == "" {
			NewError := Error{
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: "Content is missing from request headers.",
				}
				w.WriteHeader(http.StatusUnauthorized)

			NewReturnError, err := json.Marshal(NewError)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			} else {
				_, err := w.Write(NewReturnError)

				if err != nil {
					pterm.Fatal.WithFatal(true).Println(err)
				}
		}
	} else {
		NewError := Error{
			ErrorCode:    http.StatusMethodNotAllowed,
			ErrorMessage: "POST is the only accepted Method for this endpoint.",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)

		NewReturnError, err := json.Marshal(NewError)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		} else {
			_, err := w.Write(NewReturnError)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			}
		}
	}
}
