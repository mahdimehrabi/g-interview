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
	conn, err := net.Dial("tcp", s.address)
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

func (s *Socket) SendWaitJSON(data any, method string, msgID string) error {
	req := Request{
		Method: method,
		Data:   data,
		ID:     msgID,
	}
	if !s.connected {
		if err := s.Connect(); err != nil {
			return err
		}
	}
	e := json.NewEncoder(s.conn)
	if err := e.Encode(req); err != nil {
		s.connected = false
		return err
	}

	return nil
}
