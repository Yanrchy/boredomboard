package server

type MessageType int

const (
	UserConnectMSG    MessageType = 0
	UserDisconnectMSG MessageType = 1
	TextMSG           MessageType = 2
)

type Message struct {
	Type   MessageType `json:"type"`
	Sender string      `json:"sender"`
	Color  string      `json:"color"`
	Text   string      `json:"text"`
}
