package main

import (
	"github.com/mahdimehrabi/graph-interview/receiver/application/http"
	infrastructures "github.com/mahdimehrabi/graph-interview/receiver/internal/infrastructure"
)

func main() {
	env := infrastructures.NewEnv()
	env.LoadEnv()
	http.RunServer(env)
}
