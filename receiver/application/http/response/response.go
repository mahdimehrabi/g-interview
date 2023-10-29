package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Data    any      `json:"data"`
	Message string   `json:"message"`
	Status  uint16   `json:"status"`
	Errors  []string `json:"errors"`
}

func GenResponse(w http.ResponseWriter, data any, message string) {
	r, err := json.Marshal(response{
		Data:    data,
		Message: message,
	})
	if err != nil {
		log.Fatal("error")
	}

	if _, err = w.Write(r); err != nil {
		log.Fatal(err)
	}
}
