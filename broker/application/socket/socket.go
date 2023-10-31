package socket

import (
	"encoding/json"
	"fmt"
	"github.com/mahdimehrabi/graph-interview/receiver/application/socket/dto"
	infrastructures "github.com/mahdimehrabi/graph-interview/receiver/internal/infrastructure"
	"net"
)

func Handle(conn net.Conn) {
	for {
		req := dto.Request{}
		decoder := json.NewDecoder(conn)
		err := decoder.Decode(req)
		if err != nil {
			fmt.Printf("connection closed with %s", conn.RemoteAddr())
			break
		}
	}
	conn.Close()
}

func RunServer(env *infrastructures.Env) {
	ln, err := net.Listen("tcp", ":"+env.ServerPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("broker listening to port :%d 🤝" + env.ServerPort)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go Handle(conn)
	}
}
