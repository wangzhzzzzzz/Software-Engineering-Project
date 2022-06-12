package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"sync"
	"time"
)

type Course struct {
	CourseID int `json:"course_id"  gorm:"primary_key"`
	// TeacherID   string `json:"teacher_id" form:"teacher_id"`
	Name        string    `json:"name" form:"name"`
	Capacity    int       `json:"capacity" form:"capacity"`
	CapSelected int       `json:"cap_select" form:"cap_select"`
	CreateTime  time.Time `json:"create_time"`
}

func (Course) TableName() string {
	return "course"
}

type CourseDao struct {
}

var courseDao *CourseDao
var courseOnce sync.Once

func NewCourseDaoInstance() *CourseDao {
	courseOnce.Do(
		func() {
			courseDao = &CourseDao{}
		})
	return courseDao
}

func (*CourseDao) CreateCourse(course *Course) error {
	if err := db.Create(course).Error; err != nil {
		log.Println("insert course err:" + err.Error())
		return err
	}
	return nil
}

func (*CourseDao) GetCourse(CID int) (*Course, error) {
	var course Course
	err := db.Where("course_id = ?", CID).Find(&course).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Println("GetCourseByCourseID err:", err.Error())
		return nil, err
	}
	return &course, nil
}

//通过一组CID获得对应课程信息
func (*CourseDao) MGetCourse(ids []int) ([]*Course, error) {
	var courses []*Course
	err := db.Where("course_id in (?)", ids).Find(&courses).Error
	if err != nil {
		log.Println("MGetCourse err:", err.Error())
		return courses, err
	}
	return courses, nil
}
func (*CourseDao) DeleteCourse(CID int) error {
	err := db.Where("course_id = ?", CID).Delete(&Course{}).Error
	if err != nil {
		log.Println("DeleteCourse err:", err.Error())
		return err
	}
	return nil
}
func (*CourseDao) GetAllCourses(offset, limit int) ([]*Course, error) {
	var ans []*Course
	err := db.Limit(limit).Offset(offset).Find(&ans).Error
	if err != nil {
		return ans, err
	}
	return ans, nil
}
func (*CourseDao) UpdateCourseSelected(course *Course) error {
	log.Println(*course)
	err := db.Model(&Course{}).Where("course_id = ?", course.CourseID).Updates(Course{CapSelected: course.CapSelected + 1}).Error
	if err != nil {
		log.Println("UpdateCourseSelected出错")
		return err
	}
	return nil
}
