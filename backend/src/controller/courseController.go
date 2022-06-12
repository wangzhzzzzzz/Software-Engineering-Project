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
	"time"
)

//
func AddCourse(c *gin.Context) {
	AddCourseRequest := global.AddCourseRequest{}
	if err := c.ShouldBind(&AddCourseRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
		return
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
	log.Println("AddCourseRequest:", AddCourseRequest)
	//用课程名和人数限制在Course表新增一条Course记录，Bind表里新增一条记录（CID,TID）
	course := &model.Course{
		Name:        AddCourseRequest.CourseName,
		Capacity:    AddCourseRequest.Cap,
		CapSelected: 0,
		CreateTime:  time.Now().Add(8 * time.Hour),
	}
	if err := model.NewCourseDaoInstance().CreateCourse(course); err != nil {
		log.Println("Create course err:", err)
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
		return
	}
	bind := &model.Bind{
		TeacherID: memberGotten.UserID,
		CourseID:  course.CourseID,
	}
	if err := model.NewBindDaoInstance().BindCourse(bind); err != nil {
		log.Println("Bind course err:", err)
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
		return
	}

	c.JSON(http.StatusOK, global.CreateCourseResponse{Code: global.OK, Data: struct{ CourseID string }{strconv.Itoa(course.CourseID)}})
}

func TeacherGetCourse(c *gin.Context) {
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
	if user.UserType != global.Teacher {
		c.JSON(http.StatusOK, "需要教师权限")
		return
	}
	id, _ := strconv.Atoi(user.UserID)
	CIDS, err := model.NewBindDaoInstance().GetCourseIDByTeacherID(id)
	if err != nil {
		log.Println("GetCourseIDByTeacherID err:", err)
		c.JSON(http.StatusOK, "用户不存在")
		return
	}
	courses, err := model.NewCourseDaoInstance().MGetCourse(CIDS)
	if len(courses) == 0 {
		log.Println("该老师没有课程")
		c.JSON(http.StatusOK, "该老师没有课程")
		return
	}
	if err != nil {
		log.Println("批量查询时出错")
		c.JSON(http.StatusOK, "批量查询时出错")
		return
	}
	c.JSON(http.StatusOK, courses)
}

func BindCourse(c *gin.Context) {

}

func DeleteCourse(c *gin.Context) {
	DeleteCourseRequest := global.DeleteCourseRequest{}
	if err := c.ShouldBind(&DeleteCourseRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
		return
	}
	//校验courseID是否合格，然后删掉course表和bind表中与该课程有关记录
	CID, _ := strconv.Atoi(DeleteCourseRequest.CourseID)
	course, err := model.NewCourseDaoInstance().GetCourse(CID)
	if err != nil || course == nil {
		log.Println("查询了不存在的课程,err:")
		c.JSON(http.StatusOK, "课程不存在")
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	var UnBindCourseErr, DeleteCourseErr error
	go func() {
		defer wg.Done()
		err := model.NewBindDaoInstance().UnBindCourse(CID)
		if err != nil {
			UnBindCourseErr = err
			return
		}
	}()
	go func() {
		defer wg.Done()
		err := model.NewCourseDaoInstance().DeleteCourse(CID)
		if err != nil {
			DeleteCourseErr = err
			return
		}
	}()
	wg.Wait()
	if UnBindCourseErr != nil || DeleteCourseErr != nil {
		log.Println("DeleteCourse is Wrong")
		c.JSON(http.StatusOK, "删除时发生错误！")
		return
	}
	c.JSON(http.StatusOK, "删除成功")
}

func GetTeacherCourse(c *gin.Context) {

}

func ScheduleCourse(c *gin.Context) {

}

func GetCourseList(c *gin.Context) {
	//获取参数
	GetCourseListRequest := global.ListRequest{}
	if err := c.ShouldBind(&GetCourseListRequest); err != nil {
		c.JSON(http.StatusOK, "Wrong parameters!")
	}
	log.Println("GetCourseListRequest:", GetCourseListRequest)
	//读出课程，再用课程从Bind表中国查出TID，再用TID查出来教师信息，最后组合一起返回
	offset, limit := GetCourseListRequest.Offset, GetCourseListRequest.Limit
	//拿到所有的课程
	CourseList, err := model.NewCourseDaoInstance().GetAllCourses(offset, limit)
	if err != nil {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
		return
	}
	TCourseList := make([]global.TCourseAndTeacher, 0)

	for i := 0; i < len(CourseList); i++ {
		realTID, err := model.NewBindDaoInstance().GetTeacherIDByCourseID(CourseList[i].CourseID)
		if realTID == 0 || err != nil {
			if err != nil {
				log.Println("查询CID出错:", err.Error())
			}
			continue
		}
		Teacher, err := model.MemberModel.GetMemberByUserID(strconv.Itoa(realTID))
		if err != nil {
			log.Println("查询Member出错:", err.Error())
			c.JSON(http.StatusOK, global.ResponseMeta{Code: global.UnknownError})
			return
		}
		TCourseList = append(TCourseList, global.TCourseAndTeacher{
			CourseID:    strconv.Itoa(CourseList[i].CourseID),
			CourseName:  CourseList[i].Name,
			TeacherID:   Teacher.Username,
			TeacherName: Teacher.Nickname,
			Cap:         CourseList[i].Capacity,
			Selected:    CourseList[i].CapSelected,
			CreateTime:  CourseList[i].CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, global.GetStudentCourseResponse{
		Code: global.OK,
		Data: struct{ CourseList []global.TCourseAndTeacher }{CourseList: TCourseList}})
}
