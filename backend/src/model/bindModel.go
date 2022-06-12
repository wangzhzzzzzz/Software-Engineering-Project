package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type Bind struct {
	TeacherID int `json:"teacher_id" form:"teacher_id" gorm:"primary_key"`
	CourseID  int `json:"course_id" form:"course_id" gorm:"primary_key"`
}

//
// 指定Bind类对应bind表
func (Bind) TableName() string {
	return "bind"
}

type BindDao struct {
}

var bindDao *BindDao
var bindOnce sync.Once

func NewBindDaoInstance() *BindDao {
	bindOnce.Do(func() {
		bindDao = &BindDao{}
	})
	return bindDao
}

func (*BindDao) BindCourse(bind *Bind) error {
	if err := db.Create(bind).Error; err != nil {
		log.Println("bind course err:" + err.Error())
		return err
	}
	return nil
}

// UnBindCourse 取消与该课程所有的绑定记录
func (*BindDao) UnBindCourse(CID int) error {
	err := db.Where("course_id = ?", CID).Delete(&Bind{}).Error
	if err != nil {
		log.Println("DeleteCourse err:", err.Error())
		return err
	}
	return nil
}

func (*BindDao) GetTeacherIDByCourseID(CID int) (int, error) {
	var bind Bind
	err := db.Where("course_id = ?", CID).Find(&bind).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	if err != nil {
		log.Println("GetTeacherIDByCourseID err:", err.Error())
		return 0, err
	}
	return bind.TeacherID, nil
}
func (*BindDao) GetCourseIDByTeacherID(TID int) ([]int, error) {
	var bind []Bind
	var ans []int
	err := db.Where("teacher_id = ?", TID).Find(&bind).Error
	if err == gorm.ErrRecordNotFound {
		return ans, nil
	}
	if err != nil {
		log.Println("GetTeacherIDByCourseID err:", err.Error())
		return ans, err
	}
	ans = make([]int, len(bind))
	for i, b := range bind {
		ans[i] = b.CourseID
	}
	return ans, nil
}
