package dto

type Request struct {
	Method string `json:"method"`
	Data   any    `json:"data"`
}
