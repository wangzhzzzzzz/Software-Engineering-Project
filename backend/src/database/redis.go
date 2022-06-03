package database

import (
	"backend/src/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var RedisClient *redis.Pool

//
// 负责连接redis数据库和基本配置
func init() {
	//配置文件中读出来的结构体指针
	redisConf := config.GetRedisConfig()
	/*     MaxIdle:池中的最大空闲连接数。
		   MaxActive:池在给定时间分配的最大连接数。 当为零时，池中的连接数没有限制。
	       IdleTimeout:在这段时间内保持空闲后关闭连接。 如果该值为零，则空闲连接不会关闭。 应用程序应将超时设置为小于服务器超时的值。

	*/
	RedisClient = &redis.Pool{
		MaxIdle:     redisConf.MaxIdle,
		MaxActive:   redisConf.MaxActive,
		IdleTimeout: time.Duration(redisConf.TimeOut) * time.Second,
		Dial: func() (redis.Conn, error) {
			//c是连接
			c, err := redis.Dial(redisConf.Type, redisConf.Redis_Host)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			if _, err := c.Do("AUTH", redisConf.AUTH); err != nil {
				_ = c.Close()
				log.Println(err.Error())
				return nil, err
			}
			return c, nil
		},
	}
}
