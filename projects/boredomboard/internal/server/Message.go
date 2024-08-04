package server

type Message struct {
	Sender string `json:"sender"`
	Color  string `json:"color"`
	Text   string `json:"text"`
}
