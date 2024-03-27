package cmd

import (
	"chat/internal/controller/user"
	"chat/internal/middleware"
	"chat/internal/model/do/client"
	"chat/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(middleware.MiddlewareCORS)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.MiddlewareException)
				c := user.NewV1()
				group.GET("/user/login", c.Login)
				group.GET("/ws", WebSocketHandler())
			})
			s.Run()
			return nil
		},
	}
)

func WebSocketHandler() func(r *ghttp.Request) {
	room := client.NewClientRoom()
	return func(r *ghttp.Request) {
		var ctx = r.Context()
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(ctx, err)
			r.Exit()
			return
		}
		var token = r.GetQuery("token").String()
		tc := client.NewClient("", ws, &room)
		claims, err := utils.ParseTokenHs256(token)
		if err != nil {
			glog.Error(ctx, err)
			tc.WriteMessage(client.ClientMessage{
				Path: "error",
				Data: map[string]any{
					"msg": err.Error(),
				},
			})
			ws.Close()
			return
		}
		c := client.NewClient(claims.User, ws, &room)
		if room.FindByUser(c.GetUser()) != nil {
			tc.WriteMessage(client.ClientMessage{
				Path: "error",
				Data: map[string]any{
					"msg": "账号已登录",
				},
			})
			ws.Close()
			return
		}
		room.Add(c)
		c.Listening()
	}
}
