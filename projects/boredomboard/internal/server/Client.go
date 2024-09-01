package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	Channel    *CommChannel
	Connection *websocket.Conn
	Send       chan Message
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     OnCheckOrigin,
}

func (client *Client) ReadChannel() {

	for {

		select {

		case message := <-client.Send:
			client.Connection.WriteJSON(message)

		}

	}

}

func (client *Client) WriteChannel() {

	var message Message

	for {

		client.Connection.ReadJSON(&message)

		switch message.Type {

		case TextMSG:

			client.Channel.Send <- message
			message = Message{}

		}

	}

}

func OnClientConnection(channel *CommChannel, writer http.ResponseWriter, request *http.Request) {

	conn, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		log.Printf("client failed to connect | %v\n", err)
		return
	}

	client := Client{
		Channel:    channel,
		Connection: conn,
		Send:       make(chan Message),
	}

	channel.Register <- &client

	go client.WriteChannel()
	go client.ReadChannel()

	log.Printf("client connected")

}

func OnCheckOrigin(request *http.Request) bool {

	return true

}
