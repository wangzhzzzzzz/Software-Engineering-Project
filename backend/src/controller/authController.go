package controller

import (
	global "backend/src/global"
	"backend/src/model"
	"backend/src/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"time"
)

var cookiesName string = "SE-Project"

func Login(c *gin.Context) {
	loginRequest := global.LoginRequest{}
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
		return
	}
	log.Println(loginRequest)
	/*
		验证账号密码是否存在
	*/
	memberModel := model.Member{}
	memberGottenByUsername, err := memberModel.GetMemberByUsername(loginRequest.Username)
	if err != nil || memberGottenByUsername.Password != utils.Md5Encrypt(loginRequest.Password) {
		c.JSON(http.StatusOK, global.LoginResponse{Code: global.WrongPassword})
		return
	}
	if memberGottenByUsername.IsDeleted {
		c.JSON(http.StatusOK, global.LoginResponse{Code: global.UserHasDeleted})
		return
	}
	session := sessions.Default(c)
	var sessionId = getSessionId() //一个随机的uuid

	v := global.TMember{
		UserID:     strconv.Itoa(memberGottenByUsername.UserID),
		Nickname:   memberGottenByUsername.Nickname,
		Username:   memberGottenByUsername.Username,
		UserType:   memberGottenByUsername.UserType,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	session.Set(sessionId, v)
	fmt.Println(session.Get(sessionId))
	err = session.Save()
	if err != nil {
		log.Println("SaveError:", err)
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
		return
	}
	c.SetCookie(cookiesName, sessionId, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, global.LoginResponse{
		Code: http.StatusOK,
		Data: struct{ UserID string }{UserID: v.UserID},
	})
}

func Logout(c *gin.Context) {
	//先确认有没有登录
	//没有cookie的情况
	sessionId, err := c.Cookie(cookiesName)
	if err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.LoginRequired})
		return
	}
	//已经登出过了
	session := sessions.Default(c)

	if session.Get(sessionId) == nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.LoginRequired})
		return
	}
	//清除session和cookie

	session.Delete(sessionId)
	err = session.Save()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
		return
	}

	c.SetCookie(cookiesName, sessionId, -1, "/", "", false, true)
	c.JSON(http.StatusOK, global.LogoutResponse{Code: global.OK})
	//
}

func WhoAmI(c *gin.Context) {
	//先确认有没有登录
	//没有cookie的情况
	sessionId, err := c.Cookie(cookiesName)
	if err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.LoginRequired})
		return
	}
	//已经登出过了
	session := sessions.Default(c)
	Info := session.Get(sessionId) //Info is an interface{}
	if Info == nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.LoginRequired})
		return
	}
	log.Println(Info)
	user := Info.(global.TMember)
	c.JSON(http.StatusOK, global.WhoAmIResponse{Code: global.OK, Data: user})
	//
}

func getSessionId() string {
	//b := make([]byte, 32)
	//if _, err := io.ReadFull(rand.Reader, b); err != nil {
	//	return ""
	//}
	//return base64.StdEncoding.EncodeToString(b)
	//
	return uuid.NewV4().String()
}
