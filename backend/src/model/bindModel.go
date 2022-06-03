package model

import (
	"errors"
)

type Bind struct {
	TeacherID string `json:"teacher_id" form:"teacher_id" gorm:"primary_key"`
	CourseID  string `json:"course_id" form:"course_id" gorm:"primary_key"`
}

//
// 指定Bind类对应bind表
func (Bind) TableName() string {
	return "bind"
}

func (bind *Bind) BindCourse() error {

	result := db.Exec("INSERT IGNORE INTO bind(teacher_id,course_id) VALUES (?,?)", bind.TeacherID, bind.CourseID)
	if result.RowsAffected == 0 {
		return errors.New("CourseHasBound")
	}
	return nil
}
