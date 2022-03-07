package server

import (
	"backend/src/config"
	global "backend/src/global"
	router "backend/src/router"
	"encoding/gob"
	"github.com/gin-gonic/gin"
)

func Run(httpServer *gin.Engine) {
	/*
		// 生成日志，

		logFile, _ := os.Create(config.GetLogPath()) //生成该路径下的文件，注意linux和windows的文件斜杠方向不同

		gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout, os.Stdin, os.Stderr)
		// 设置日志格式
		httpServer.Use(gin.LoggerWithFormatter(config.GetLogFormat))
		httpServer.Use(gin.Recovery())
	*/

	//设置session,session解析TMember
	gob.Register(global.TMember{})
	httpServer.Use(global.GetSession())

	// 注册路由
	router.RegisterRouter(httpServer)

	serverError := httpServer.Run(config.GetServerConfig().HTTP_HOST + ":" + config.GetServerConfig().HTTP_PORT)

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}

}
