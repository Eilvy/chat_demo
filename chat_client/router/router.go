package router

import (
	"github.com/gin-gonic/gin"
	"go_code/chat_demo/chat_client/service"
)

func InitRouter() {
	r := gin.Default()

	u := r.Group("/user")
	{
		u.POST("/register", service.Register)
		u.GET("/login", service.Login)
	}
	c := r.Group("/chat")
	{
		c.GET("/:roomId", service.WsChat)
	}

	err := r.Run("8080")
	if err != nil {
		return
	}
}
