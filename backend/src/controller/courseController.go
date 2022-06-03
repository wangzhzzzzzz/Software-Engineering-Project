package controller

import (
	global "backend/src/global"
	"backend/src/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//
func AddCourse(c *gin.Context) {
	AddCourseRequest := global.AddCourseRequest{}
	if err := c.ShouldBind(&AddCourseRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
	}
	//先判断教工号和教师名称是否合法
	memberGotten, err := model.MemberModel.GetMemberByUsername(AddCourseRequest.Username)
	if err != nil || memberGotten.UserID == 0 {
		// 教师不存在
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UserNotExisted})
		return
	}
	if memberGotten.IsDeleted {
		// 教师已经被删除
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UserHasDeleted})
		return
	}
	if memberGotten.UserType != global.Teacher || memberGotten.Nickname != AddCourseRequest.TeacherName {
		// 该成员不是教师
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.CourseNotAvailable})
		return
	}
	//开启事务完成，用课程名和人数限制在Course表新增一条Course记录，Bind表里新增一条记录
	/*
		c.JSON(http.StatusOK, global.CreateCourseResponse{Code: global.OK, Data: struct{ CourseID string }{CID}})

	*/
}

func GetCourse(c *gin.Context) {

}

func BindCourse(c *gin.Context) {

}

func UnbindCourse(c *gin.Context) {

}

func GetTeacherCourse(c *gin.Context) {

}

func ScheduleCourse(c *gin.Context) {

}

func GetCourseList(c *gin.Context) {
	GetCourseListRequest := global.ListRequest{}
	if err := c.ShouldBind(&GetCourseListRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
	}
	log.Println("request:", GetCourseListRequest)
	courseModel := model.Course{}
	offset, limit := GetCourseListRequest.Offset, GetCourseListRequest.Limit
	CourseList, err := courseModel.GetAllCourses(offset, limit)
	if err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
		return
	}
	TCourseList := make([]global.TCourseAndTeacher, 0)
	for i := 0; i < len(CourseList); i++ {
		TCourseList = append(TCourseList, global.TCourseAndTeacher{
			CourseID:    strconv.Itoa(CourseList[i].CourseID),
			CourseName:  CourseList[i].Name,
			TeacherID:   "暂定",
			TeacherName: "暂定",
			Cap:         CourseList[i].Capacity,
			Selected:    CourseList[i].CapSelected,
			CreateTime:  CourseList[i].CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, global.GetStudentCourseResponse{
		Code: global.OK,
		Data: struct{ CourseList []global.TCourseAndTeacher }{CourseList: TCourseList}})
}
