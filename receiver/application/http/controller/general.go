package controller

import (
	"net/http"

	"github.com/mahdimehrabi/graph-interview/receiver/application/http/response"
)

type General struct{}

func NewGeneral() *General {
	return &General{}
}

func (g General) Ping(w http.ResponseWriter, _ *http.Request) {
	response.GenResponse(w, map[string]string{"ping": "pong🏓"}, "")
}
