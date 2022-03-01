package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"time"
)

type MySQLInfo struct {
	TYPE         string
	USER         string
	PASSWORD     string
	DB_HOST      string
	DB_PORT      string
	DB_NAME      string
	CHARSET      string
	ParseTime    string
	MaxIdleConns int
	MaxOpenConns int
}

type ServerInfo struct {
	HTTP_PORT string
	HTTP_HOST string
	MODE      string
}

type RedisInfo struct {
	Redis_Host string
	Type       string
	MaxIdle    int
	MaxActive  int
	TimeOut    int
}

func GetServerConfig() *ServerInfo {
	cfg, err := ini.Load("configFile/config.ini")
	if err != nil {
		log.Printf("Fail to read file: %v \n", err)
		os.Exit(1)
	}
	d := new(ServerInfo)
	_ = cfg.Section("server").MapTo(d)
	return d
}

func GetMySQLConfig() *MySQLInfo {
	cfg, err := ini.Load("configFile/config.ini")
	if err != nil {
		log.Printf("Fail to read file: %v \n", err)
		os.Exit(1)
	}
	d := new(MySQLInfo)
	_ = cfg.Section("mysql").MapTo(d)
	return d
}

func GetRedisConfig() *RedisInfo {
	cfg, err := ini.Load("configFile/config.ini")
	if err != nil {
		log.Printf("Fail to read file: %v \n", err)
		os.Exit(1)
	}
	d := new(RedisInfo)
	_ = cfg.Section("redis").MapTo(d)
	return d
}

// GetLogPath 生成文件路径字符串
func GetLogPath() string {
	timeObj := time.Now()
	datetime := timeObj.Format("2006-01-02-15-04-05")
	return "C:\\Users\\86139\\Desktop\\grade-3\\软件工程导论\\project\\" + datetime + ".log"
}

// GetLogFormat 设置日志格式
func GetLogFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
