package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	Timeout      = 30 * time.Second
	PongTimeout  = 30 * time.Second
	PingInterval = 20 * time.Second
)

type Client struct {
	Channel    *CommChannel
	Connection *websocket.Conn
	Send       chan Message
	Control    chan int
	Disconnect chan *Client

	Username string
	Color    string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     OnCheckOrigin,
}

func (client *Client) ReadChannel() {

	defer client.Connection.Close()

	for {
		select {

		case message := <-client.Send:
			client.Connection.WriteJSON(message)

		case message := <-client.Control:
			client.Connection.WriteControl(message, nil, time.Now().Add(Timeout))

		case <-client.Disconnect:
			return

		}
	}

}

func (client *Client) WriteChannel() {

	defer client.Connection.Close()

	/*************************************************************************************
	 * Sets the cleint's pong handler, failure to meet the set deadline results in       *
	 * error in subsequent reads in the message loop, the loop should handle how to deal *
	 * with that.                                                                        *
	 *************************************************************************************/
	client.Connection.SetReadDeadline(time.Now().Add(PongTimeout))
	client.Connection.SetPongHandler(func(appData string) error {

		log.Printf("pong msg received\n")
		err := client.Connection.SetReadDeadline(time.Now().Add(PongTimeout))

		return err
	})

	var message Message
	for {

		err := client.Connection.ReadJSON(&message)
		if err != nil {

			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v \n", err)
				log.Printf("terminating connection %v", client.Username)
				client.Channel.Unregister <- client
			}

			return
		}

		switch message.Type {

		case UserConnectMSG:
			client.Username = message.Sender
			client.Color = message.Color

			client.Channel.BroadcastNewUser <- message
			message = Message{}

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
		Control:    make(chan int),
		Disconnect: make(chan *Client),
	}

	channel.Register <- &client

	go client.WriteChannel()
	go client.ReadChannel()

	log.Printf("client connected")

}

func OnCheckOrigin(request *http.Request) bool {

	return true

}
