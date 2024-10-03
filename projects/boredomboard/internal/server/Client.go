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

/*
 * Client is the representation of the "client" on the server side
 * It has two main methods for interacting with the CommChannel.
 *
 * func (client *Client) ReadChannel();
 * 	Means the client is reading the chat state on the server side.
 * 	If the client is to receive messages, it's through this method.
 *
 * func (client *Client) WriteChannel();
 * 	Means the client is writing to the chat state on the server side.
 *
 * The messages that are sent and received by the client can have different meanings,
 * use MessageType enums to evaluate said messages.
 */

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

		// uncomment this for testing
		// log.Printf("server: pong msg received from %v\n", client.Username)
		err := client.Connection.SetReadDeadline(time.Now().Add(PongTimeout))

		return err
	})

	var message Message
	for {

		/*
		 * When ReadJSON throws an error, the CommChannel should disconnect the client.
		 *	If the client fails to respond to control frames, it will throw an error.
		 *
		 *	The function will also throw an error if the client sends a disconnect control frame,
		 *	this is the default close handler of gorilla websocket, all subsequent reads will throw
		 *	an error if we receive a conntrol frame for disconnection from the client. When that happens
		 *	CommChannel should disconnect the client.
		 */

		err := client.Connection.ReadJSON(&message)
		if err != nil {

			//if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
			//}

			log.Printf("server: %v\n", err)
			log.Printf("server: terminating connection with %v\n", client.Username)
			client.Channel.Unregister <- client
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
		log.Printf("server: client failed to connect, %v\n", err)
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

	log.Printf("server: new client connected\n")

}

func OnCheckOrigin(request *http.Request) bool {

	return true

}
