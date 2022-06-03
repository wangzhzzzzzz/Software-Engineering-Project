package types

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var cookiesName string = "SE-Project"

//
func BackendAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println(c.FullPath())

		session := sessions.Default(c)
		sessionId, err := c.Cookie(cookiesName)

		if err != nil {
			log.Println("第一个err有问题")
			c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
			c.Abort()
			return
		}

		v := session.Get(sessionId)
		if v == nil {
			log.Println("第二个err有问题")
			c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
			c.Abort()
			return
		}
		// log.Println(v)
		user := v.(TMember)

		if user.UserType != Admin {
			log.Println("第三个err有问题")
			c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
			c.Abort()
			return
		}
		c.Next()
	}
}
