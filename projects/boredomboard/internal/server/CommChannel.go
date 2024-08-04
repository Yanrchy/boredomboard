package server

import "log"

type CommChannel struct {
	Transcript []Message
	Lobby      map[*Client]bool

	Register   chan *Client
	Unregister chan *Client

	Send chan Message
}

func NewChannel() *CommChannel {

	return &CommChannel{
		Transcript: make([]Message, 0, 1024),
		Lobby:      make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Send:       make(chan Message),
	}

}

func (channel *CommChannel) Serve() {

	for {

		select {

		case client := <-channel.Register:
			channel.Lobby[client] = true

			for i := 0; i < len(channel.Transcript); i++ {
				client.Send <- channel.Transcript[i]
			}

			//case client := <-channel.Unregister:

		case message := <-channel.Send:
			channel.Transcript = append(channel.Transcript, message)

			for client := range channel.Lobby {

				client.Send <- message

			}

			log.Printf("%v: %v", message.Sender, message.Text)

		}

	}

}
