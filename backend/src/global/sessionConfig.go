package types

import (
	"backend/src/database"

	"github.com/gin-contrib/sessions"
	redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//
func GetSession() gin.HandlerFunc {
	// NewStoreWithPool 使用传入的 *redis.Pool 实例化一个 RediStore。
	// database.RedisClient就是连接的*redis.Pool
	store, _ := redis.NewStoreWithPool(database.RedisClient, []byte("passord"))
	return sessions.Sessions("MySession", store)
}
