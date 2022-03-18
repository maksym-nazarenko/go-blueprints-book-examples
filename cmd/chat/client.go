package main

import (
	"encoding/json"
	"log"
	"time"

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
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			log.Printf("error reading WS message %v", err)
			return
		}
		msg.When = time.Now()
		name := "Guest"
		if nameIntf, ok := c.userData["name"]; ok {
			name = nameIntf.(string)
		}
		msg.Name = name
		if avatarURL, ok := c.userData["avatar_url"]; ok {
			msg.AvatarURL = avatarURL.(string)
		}

		c.room.forward <- msg
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
