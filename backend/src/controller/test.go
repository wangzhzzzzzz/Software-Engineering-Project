package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type testRequest struct {
	Str string `form:"gg" binding:"required"`
}

type testResponse struct {
	Str string `json:"123" form:"nickname"`
}

func Ping(c *gin.Context) {
	request := testRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, "搞错啦")
		return
	}

	c.JSON(200, request)
	c.JSON(200, testResponse{Str: "hello world dongshujie"})
}
