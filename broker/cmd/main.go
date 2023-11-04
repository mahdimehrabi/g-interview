package main

import (
	"github.com/mahdimehrabi/graph-interview/broker/application/socket"
	"github.com/mahdimehrabi/graph-interview/broker/internal"
	infrastructures "github.com/mahdimehrabi/graph-interview/broker/internal/infrastructure"
)

func main() {
	env := infrastructures.NewEnv()
	env.LoadEnv()
	internal.SetupDPI(env)
	socket.RunServer(env)
}
