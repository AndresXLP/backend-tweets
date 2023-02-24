package entity

type Message struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
