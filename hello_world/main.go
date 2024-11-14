package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"go_code/hello_world/api"
	"go_code/hello_world/utils"
)

func main() {
	utils.DBInit()
	fmt.Println("hello world")
	r := *server.Default(server.WithHostPorts(":3000"))
	r.POST("/login", api.WXLogin)
	r.Spin()
}
