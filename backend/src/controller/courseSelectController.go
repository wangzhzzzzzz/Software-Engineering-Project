package controller

import (
	global "backend/src/global"
	"backend/src/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//
func BookCourse(c *gin.Context) {
	bookCourseRequest := global.BookCourseRequest{}
	if err := c.ShouldBind(&bookCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.BookCourseResponse{Code: global.UnknownError})
		return
	}
	//校验学生身份
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
	user := Info.(global.TMember)
	if user.UserType != global.Student {
		c.JSON(http.StatusOK, "学生才能选课")
		return
	}
	//检验课程信息,判断该课程的容量是否还足够
	course, err := model.NewCourseDaoInstance().GetCourse(bookCourseRequest.CourseID)
	if course == nil {
		c.JSON(http.StatusOK, "课程不存在")
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, "选课异常出错")
		return
	}
	if course.CapSelected == course.Capacity {
		c.JSON(http.StatusOK, "课程容量已满")
		return
	}
	//校验是否选过该门课程
	userID, _ := strconv.Atoi(user.UserID)
	choice := model.Choice{
		StudentID: userID,
		CourseID:  bookCourseRequest.CourseID,
	}
	if model.NewChoiceDaoInstance().IsExisted(&choice) {
		c.JSON(http.StatusOK, "已经选过该课程")
		return
	}
	//选课，Choice表中加一条记录，更新该课程selected+1
	var wg sync.WaitGroup
	wg.Add(2)
	var SelectCourseErr, UpdateCourseSelectedErr error
	go func() {
		defer wg.Done()
		err := model.NewChoiceDaoInstance().SelectCourse(&choice)
		if err != nil {
			SelectCourseErr = err
			return
		}
	}()
	go func() {
		defer wg.Done()
		err := model.NewCourseDaoInstance().UpdateCourseSelected(course)
		if err != nil {
			UpdateCourseSelectedErr = err
			return
		}
	}()
	wg.Wait()
	if SelectCourseErr != nil || UpdateCourseSelectedErr != nil {
		c.JSON(http.StatusOK, "更新选课信息出错")
		return
	}
	c.JSON(http.StatusOK, global.BookCourseResponse{Code: global.OK})
	return
}

func GetStudentCourse(c *gin.Context) {
	//校验学生身份
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
	user := Info.(global.TMember)
	log.Println("要查课表的学生:", user)
	if user.UserType != global.Student {
		c.JSON(http.StatusOK, "学生才能查看选课信息")
		return
	}
	userID, _ := strconv.Atoi(user.UserID)
	CIDS, err := model.NewChoiceDaoInstance().QueryCourseSelected(userID)
	if err != nil {
		c.JSON(http.StatusOK, "查询课表时出错")
		return
	}
	courses, err := model.NewCourseDaoInstance().MGetCourse(CIDS)
	if err != nil {
		c.JSON(http.StatusOK, "查询课表时出错")
		return
	}
	c.JSON(http.StatusOK, courses)
	return
}
