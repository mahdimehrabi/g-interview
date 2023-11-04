package socket

import (
	"encoding/json"
	"net"

	"github.com/mahdimehrabi/graph-interview/broker/application/socket/dto"
	infrastructures "github.com/mahdimehrabi/graph-interview/broker/internal/infrastructure"
	"github.com/rs/zerolog/log"
)

func Handle(conn net.Conn) {
	for {
		req := dto.Request{}
		decoder := json.NewDecoder(conn)
		err := decoder.Decode(&req)
		if err != nil {
			log.Print("error extracting data from socket", err)
			log.Printf("connection closed with %s", conn.RemoteAddr())
			break
		}
		HandleRoute(conn, req)
	}
	conn.Close()
}

func RunServer(env *infrastructures.Env) {
	ln, err := net.Listen("tcp", ":"+env.ServerPort)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("broker listening to port :%s 🤝", env.ServerPort)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		go Handle(conn)
	}
}
