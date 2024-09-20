package server

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type CommChannel struct {
	Transcript []Message
	Lobby      map[*Client]bool

	Register   chan *Client
	Unregister chan *Client

	Send             chan Message
	BroadcastNewUser chan Message
}

func NewChannel() *CommChannel {

	return &CommChannel{
		Transcript:       make([]Message, 0, 1024),
		Lobby:            make(map[*Client]bool),
		Register:         make(chan *Client),
		Unregister:       make(chan *Client),
		Send:             make(chan Message),
		BroadcastNewUser: make(chan Message),
	}

}

func (channel *CommChannel) Serve() {

	/*********** Ticker for when to ping the client ***********/
	/**********************************************************/
	ticker := time.NewTicker(PingInterval)
	defer ticker.Stop()

	tickChannel := make(chan time.Time)

	go func() {
		for tick := range ticker.C {
			tickChannel <- tick
		}
	}()
	/**********************************************************/

	for {

		select {

		case client := <-channel.Register:
			channel.Lobby[client] = true

			for i := 0; i < len(channel.Transcript); i++ {
				client.Send <- channel.Transcript[i]
			}

			for clientConnected := range channel.Lobby {
				client.Send <- Message{
					Type:   UserConnectMSG,
					Sender: clientConnected.Username,
					Color:  clientConnected.Color,
					Text:   "",
				}
			}

			//case client := <-channel.Unregister:

		case message := <-channel.Send:
			channel.Transcript = append(channel.Transcript, message)

			for client := range channel.Lobby {

				client.Send <- message

			}

			log.Printf("%v: %v", message.Sender, message.Text)

		case message := <-channel.BroadcastNewUser:

			for client := range channel.Lobby {

				client.Send <- message

			}

		case <-tickChannel:
			for client := range channel.Lobby {

				client.Control <- websocket.PingMessage

			}

		}

	}

}
