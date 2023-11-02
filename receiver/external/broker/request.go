package broker

type Request struct {
	Method string `json:"method"`
	Data   any    `json:"data"`
	ID     string `json:"ID"`
}
