package route

import (
	"net/http"

	"github.com/mahdimehrabi/graph-interview/receiver/application/http/controller"
)

type General struct {
	mux               *http.ServeMux
	generalController *controller.General
}

func NewGeneral(mux *http.ServeMux) General {
	return General{
		mux:               mux,
		generalController: controller.NewGeneral(),
	}
}

func (g General) Handle() {
	http.HandleFunc("/api/ping", g.generalController.Ping)
}
