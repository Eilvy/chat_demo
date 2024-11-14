package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go_code/chat_demo/chat_client/resps"
	"log"
	"net/http"
	"sync"
)

var (
	mux sync.Mutex
)

type client struct {
	conn     *websocket.Conn
	username string
	send     chan []byte
}

// Upgrade 升级为websocket的结构体
var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsChat 聊天室主程序
func WsChat(ctx *gin.Context) {
	conn, err := Upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("websocket upgrade error : ", err)
		resps.InternalErr(ctx)
		return
	}
}
