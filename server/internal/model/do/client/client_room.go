package client

import (
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

func NewClientRoom() ClientRoom {
	return ClientRoom{
		clients: map[string]*Client{},
	}
}

type ClientRoom struct {
	clients map[string]*Client
}

func (s *ClientRoom) Add(client *Client) {
	s.clients[client.GetUser()] = client
}

func (s *ClientRoom) RemoveByUser(user string) {
	delete(s.clients, user)
}

func (s *ClientRoom) FindByUser(user string) *Client {
	for key, client := range s.clients {
		if key == user {
			return client
		}
	}
	return nil
}

type ClientMessage struct {
	Path string         `json:"path" c:"path" v:"required"`
	Data map[string]any `json:"data" c:"data" v:"required"`
}

func (s *ClientRoom) SendMsgToAll(c *Client, m ClientMessage) int {
	var n = 0
	for _, client := range s.clients {
		if client != nil {
			if err := client.WriteMessage(m); err != nil {
				glog.Error(context.TODO(), c.user, "->", client.user, "==", m)
			} else {
				n++
			}
		}
	}
	return n
}
