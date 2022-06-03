package model

import (
	"backend/src/database"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
	"strconv"
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

func (course *Course) CreateCourse() (string, error) {
	log.Println(course)
	newcourse := *course
	err := db.Create(&newcourse).Error
	if err != nil {
		log.Println(err)
		return "", err
	}
	return strconv.Itoa(newcourse.CourseID), nil
}

func (course *Course) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	return scope.SetColumn("course_id", uuid)
}

func GetCourse(course_id string) (Course, error) {
	var result Course
	err := database.MySqlDb.First(&Course{}, "course_id = ?", course_id).Scan(&result).Error
	return result, err
}
func (course *Course) GetAllCourses(offset, limit int) ([]Course, error) {
	var ans []Course
	err := db.Limit(limit).Offset(offset).Find(&ans).Error
	if err != nil {
		return ans, err
	}
	return ans, nil
}
