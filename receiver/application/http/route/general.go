package route

import (
	"net/http"

	"github.com/mahdimehrabi/graph-interview/receiver/application/http/controller/general"
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

func (g *General) Handle() {
	g.mux.HandleFunc("/api/ping", g.generalController.Ping)
}
