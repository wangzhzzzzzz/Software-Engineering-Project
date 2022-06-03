package controller

import (
	"backend/src/database"
	global "backend/src/global"
	"backend/src/model"
	"backend/src/utils"
	"backend/src/validate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

//
func CreateMember(c *gin.Context) {
	//接收参数
	createMemberRequest := global.CreateMemberRequest{}
	if err := c.ShouldBind(&createMemberRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
		return
	}
	log.Println(createMemberRequest)
	//验证参数
	if !validate.IsValidCreateMember(&createMemberRequest) {
		c.JSON(http.StatusOK, "Wrong parameters!")
		return
	}
	//判断是不是管理员在中间件中实现过
	memberModel := model.Member{
		Nickname:   createMemberRequest.Nickname,
		Username:   createMemberRequest.Username,
		Password:   utils.Md5Encrypt(createMemberRequest.Password),
		UserType:   createMemberRequest.UserType,
		CreateTime: time.Now(),
	}
	Id, err := memberModel.CreateMember() //创建一个用户并返回用户ID号
	if err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UserHasExisted})
		return
	}
	//把创建的用户ID存入redis
	if createMemberRequest.UserType == global.Student {
		rc := database.RedisClient.Get()
		model.AddStudentID(Id, rc)
		rc.Close()
	}
	c.JSON(http.StatusOK, global.CreateMemberResponse{Code: global.OK, Data: struct{ UserID string }{Id}})
}

func GetMember(c *gin.Context) {
	//接收参数
	getMemberRequest := global.GetMemberRequest{}
	if err := c.ShouldBind(&getMemberRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
		return
	}
	memberModel := model.Member{}
	log.Println(getMemberRequest)
	memberGotten, err := memberModel.GetMemberByUserID(getMemberRequest.UserID)
	if err != nil || memberGotten.UserID == 0 {
		// 用户不存在
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UserNotExisted})
		return
	}
	if memberGotten.IsDeleted {
		// 用户已经被删除
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UserHasDeleted})
		return
	}
	c.JSON(http.StatusOK, global.GetMemberResponse{
		Code: global.OK,
		Data: global.TMember{
			UserID:     strconv.Itoa(memberGotten.UserID),
			Nickname:   memberGotten.Nickname,
			Username:   memberGotten.Username,
			CreateTime: memberGotten.CreateTime.Format("2006-01-02 15:04:05"),
			IsDeleted:  "未注销",
		},
	})
}

func GetStudentList(c *gin.Context) {
	GetMemberListRequest := global.ListRequest{}
	if err := c.ShouldBind(&GetMemberListRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
	}
	memberModel := model.Member{}
	log.Println(GetMemberListRequest)
	offset, limit := GetMemberListRequest.Offset, GetMemberListRequest.Limit
	MemberList, err := memberModel.GetAllMembers(offset, limit)
	if err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
		return
	}

	TMemberList := make([]global.TMember, 0)
	for _, v := range MemberList {
		if v.IsDeleted || v.UserType != global.Student {
			continue
		}
		TMemberList = append(TMemberList, global.TMember{
			UserID:     strconv.Itoa(v.UserID),
			Nickname:   v.Nickname,
			Username:   v.Username,
			CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
			UserType:   v.UserType,
		})

	}
	// 返回
	c.JSON(http.StatusOK, global.GetMemberListResponse{
		Code: global.OK,
		Data: struct{ MemberList []global.TMember }{MemberList: TMemberList}})
}

func GetTeacherList(c *gin.Context) {
	GetMemberListRequest := global.ListRequest{}
	if err := c.ShouldBind(&GetMemberListRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
	}
	memberModel := model.Member{}
	log.Println(GetMemberListRequest)
	offset, limit := GetMemberListRequest.Offset, GetMemberListRequest.Limit
	MemberList, err := memberModel.GetAllMembers(offset, limit)
	if err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
		return
	}

	TMemberList := make([]global.TMember, 0)
	for _, v := range MemberList {
		if v.IsDeleted || v.UserType != global.Teacher {
			continue
		}
		TMemberList = append(TMemberList, global.TMember{
			UserID:     strconv.Itoa(v.UserID),
			Nickname:   v.Nickname,
			Username:   v.Username,
			CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
			UserType:   v.UserType,
		})

	}
	// 返回
	c.JSON(http.StatusOK, global.GetMemberListResponse{
		Code: global.OK,
		Data: struct{ MemberList []global.TMember }{MemberList: TMemberList}})
}

func UpdateMember(c *gin.Context) {
	// 用于定义接受哪些请求的参数
	updateMemberRequest := global.UpdateMemberRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&updateMemberRequest); err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.ParamInvalid})
		return
	}

	log.Println(updateMemberRequest)

	err := model.UpdateMember(updateMemberRequest.UserID, updateMemberRequest.Nickname, updateMemberRequest.Username)

	if err == nil {
		// 成功更新用户昵称
		c.JSON(http.StatusOK, global.UpdateMemberResponse{Code: global.OK})
		return
	}

	if err.Error() == "未找到该用户" {
		// 用户不存在
		c.JSON(http.StatusOK, global.UpdateMemberResponse{Code: global.UserNotExisted})
		return
	} else if err.Error() == "用户已删除" {
		// 用户已删除
		c.JSON(http.StatusOK, global.UpdateMemberResponse{Code: global.UserHasDeleted})
		return
	}
}

func DeleteMember(c *gin.Context) {
	// 用于定义接受哪些请求的参数
	deleteMemberRequest := global.DeleteMemberRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&deleteMemberRequest); err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.ParamInvalid})
		return
	}

	log.Println(deleteMemberRequest)

	err := model.DeleteMember(deleteMemberRequest.UserID)
	if err != nil {
		if err.Error() == "未找到该用户" {
			// 用户不存在
			log.Println(err)
			c.JSON(http.StatusOK, global.DeleteMemberResponse{Code: global.UserNotExisted})
			return
		} else if err.Error() == "用户已删除" {
			// 用户已删除
			log.Println(err)
			c.JSON(http.StatusOK, global.DeleteMemberResponse{Code: global.UserHasDeleted})
			return
		} else {
			log.Println(err)
			c.JSON(http.StatusOK, global.DeleteMemberResponse{Code: global.UnknownError})
		}
	}

	// 成功删除用户
	rc := database.RedisClient.Get()
	model.RemoveStudentID(deleteMemberRequest.UserID, rc)
	rc.Close()
	c.JSON(http.StatusOK, global.DeleteMemberResponse{Code: global.OK})
	return

}
