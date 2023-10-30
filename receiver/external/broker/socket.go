package broker

import (
	"encoding/json"
	"net"
)

type Socket struct {
	address   string
	conn      net.Conn
	connected bool
}

func NewSocket(address string) *Socket {
	return &Socket{
		address: address,
	}
}

func (s *Socket) Connect() error {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		return err
	}
	s.conn = conn
	s.connected = true
	return nil
}

func (s *Socket) Disconnect() error {
	s.connected = false
	return s.conn.Close()
}

func (s *Socket) SendJSON(data any) error {
	e := json.NewEncoder(s.conn)
	if err := e.Encode(data); err != nil {
		s.connected = false
		return err
	}
	return nil
}
