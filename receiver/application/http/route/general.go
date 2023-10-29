package route

import (
	"github.com/mahdimehrabi/graph-interview/receiver/application/http/controller/general"
	"net/http"
)

type General struct {
	mux               *http.ServeMux
	generalController *general.General
}

func NewGeneral(mux *http.ServeMux) General {
	return General{
		mux:               mux,
		generalController: general.NewGeneral(),
	}
}

func (g General) Handle() {
	http.HandleFunc("/api/ping", g.generalController.Ping)
}
