package main

import (
	"github.com/mahdimehrabi/graph-interview/receiver/application/http"
	"github.com/mahdimehrabi/graph-interview/receiver/internal"
	infrastructures "github.com/mahdimehrabi/graph-interview/receiver/internal/infrastructure"
)

func main() {
	env := infrastructures.NewEnv()
	env.LoadEnv()
	internal.SetupDPI(env)
	http.RunServer(env)
}
