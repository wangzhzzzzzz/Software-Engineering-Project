package main

import (
	"backend/src/server"
	"github.com/gin-gonic/gin"
)

func main() {
	httpServer := gin.Default() //httpServer是路由引擎，提供网络处理能力
	server.Run(httpServer)      //传入server包中的Run函数
}
