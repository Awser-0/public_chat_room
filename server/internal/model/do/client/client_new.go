package client

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func NewClient(user string, ws *ghttp.WebSocket, room *ClientRoom) *Client {
	return &Client{
		user: user,
		ws:   ws,
		room: room,
	}
}

type Client struct {
	user string
	ws   *ghttp.WebSocket
	room *ClientRoom
}

func (c *Client) GetUser() string {
	return c.user
}

func (c *Client) WriteMessage(message ClientMessage) error {
	return c.ws.WriteJSON(message)
}

func jsonStrToMessage(jsonStr string) (ClientMessage, error) {
	var msg ClientMessage
	var gj *gjson.Json
	var err error
	if gj, err = gjson.DecodeToJson(jsonStr); err != nil {
		return msg, err
	}
	if err = gj.Scan(&msg); err != nil {
		return msg, err
	}
	return msg, nil
}

func (c *Client) Listening() {
	for {
		_, jsonStrB, err := c.ws.ReadMessage()
		if err != nil {
			c.ws.Close()
			c.room.RemoveByUser(c.GetUser())
			glog.Error(context.TODO(), "接受消息失败", err)
			break
		}

		msg, err := jsonStrToMessage(string(jsonStrB))
		if err != nil {
			fmt.Println("数据转化失败：", err.Error())
			glog.Error(context.TODO(), "数据转化失败：", err)
			continue
		}

		if msg.Path == "/chat" {
			if c.room != nil {
				sendMsg := ClientMessage{
					Path: "/chat",
					Data: map[string]any{
						"text":        msg.Data["text"],
						"user":        c.user,
						"create_time": time.Now(),
					},
				}
				c.room.SendMsgToAll(c, sendMsg)
			}
		}
	}
}
