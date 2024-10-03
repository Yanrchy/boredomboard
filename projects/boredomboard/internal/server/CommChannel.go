package server

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

/*
 * CommChannel represents the actual state of the chat application.
 * It's reponsible for keeping track of the message Transcript
 * and all the clients connected to the lobby.
 *
 * Clients may read and write to it, since this is the true state of the chat,
 * the channel should send a signal to the connected client's frontend app to
 * update the state on client side.
 */
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

		case client := <-channel.Unregister:
			for connectedClient := range channel.Lobby {
				connectedClient.Send <- Message{
					Type:   UserDisconnectMSG,
					Sender: client.Username,
					Color:  client.Color,
					Text:   "",
				}
			}

			client.Disconnect <- client
			delete(channel.Lobby, client)

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
