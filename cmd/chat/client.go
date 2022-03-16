package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			log.Printf("error reading WS message %v", err)
			return
		}
		var receivedMsg struct {
			Message string
		}

		if err := json.Unmarshal(msg, &receivedMsg); err != nil {
			log.Printf("error unmarshalling WS message %v", err)
			return
		}
		log.Printf("error reading WS message %v", err)
		name := "Guest"
		if nameIntf, ok := c.userData["name"]; ok {
			name = nameIntf.(string)
		}
		readMsg := &message{Message: receivedMsg.Message, Name: name}
		c.room.tracer.Trace(fmt.Sprintf("message: %+v", readMsg))

		c.room.forward <- readMsg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		data, err := json.Marshal(msg)
		if err != nil {
			log.Printf("message marchalling failed: %v", err)
			continue
		}
		err = c.socket.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			return
		}
	}
}
