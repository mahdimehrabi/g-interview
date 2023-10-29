package http

import (
	"log"
	"net/http"
	"time"

	"github.com/mahdimehrabi/graph-interview/receiver/application/http/route"
	infrastructures "github.com/mahdimehrabi/graph-interview/receiver/internal/infrastructure"
)

func RunServer(env *infrastructures.Env) {
	mx := http.NewServeMux()
	general := route.NewGeneral(mx)
	general.Handle()
	server := &http.Server{
		Addr:              ":" + env.ServerPort,
		Handler:           mx,
		ReadHeaderTimeout: time.Second * 5, // prevent Slow-loris attack
	}
	log.Printf("running http REST API on port \":%s\" 🏁", env.ServerPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
