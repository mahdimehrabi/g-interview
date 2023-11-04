package main

import (
	"github.com/mahdimehrabi/graph-interview/destination/application/socket"
	"github.com/mahdimehrabi/graph-interview/destination/internal"
	infrastructures "github.com/mahdimehrabi/graph-interview/destination/internal/infrastructure"
)

func main() {
	env := infrastructures.NewEnv()
	env.LoadEnv()
	internal.SetupDPI()
	socket.RunServer(env)
}
