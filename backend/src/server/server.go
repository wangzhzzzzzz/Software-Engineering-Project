package server

import (
	"io"
	"os"
	"project/src/config"
	router "project/src/router"

	"github.com/gin-gonic/gin"
)

func Run(httpServer *gin.Engine) {

	// 生成日志，

	logFile, _ := os.Create(config.GetLogPath()) //生成该路径下的文件，注意linux和windows的文件斜杠方向不同

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout, os.Stdin, os.Stderr)
	// 设置日志格式
	httpServer.Use(gin.LoggerWithFormatter(config.GetLogFormat))
	httpServer.Use(gin.Recovery())

	// 注册路由
	router.RegisterRouter(httpServer)

	serverError := httpServer.Run(config.GetServerConfig().HTTP_HOST + ":" + config.GetServerConfig().HTTP_PORT)

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}

}
