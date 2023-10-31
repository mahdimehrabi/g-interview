package response

import (
	"encoding/json"
	"io"
	"log"
)

type response struct {
	Data    any      `json:"data"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func GenResponse(w io.Writer, data any, message string, errors []string) {
	r, err := json.Marshal(response{
		Data:    data,
		Message: message,
		Errors:  errors,
	})
	if err != nil {
		log.Fatal("error")
	}

	if _, err = w.Write(r); err != nil {
		log.Fatal(err)
	}
}

func SuccessResponse(w io.Writer, data any, message string) {
	GenResponse(w, data, message, []string{})
}

func InternalErrorResponse(w io.Writer) {
	GenResponse(w, nil, "internal server error", []string{})
}

func BadRequestErrorResponse(w io.Writer) {
	GenResponse(w, nil, "bad request", []string{})
}
