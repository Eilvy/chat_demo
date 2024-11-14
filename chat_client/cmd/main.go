package main

import (
	"go_code/chat_demo/chat_client/dao"
	"go_code/chat_demo/chat_client/router"
)

func main() {
	dao.Init()
	dao.InitRedis()
	router.InitRouter()
}
