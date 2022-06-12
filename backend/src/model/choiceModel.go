package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type Choice struct {
	StudentID int `json:"student_id" form:"student_id" gorm:"primary_key"`
	CourseID  int `json:"course_id" form:"course_id" gorm:"primary_key"`
}

func (Choice) TableName() string {
	return "choice"
}

type ChoiceDao struct {
}

var choiceDao *ChoiceDao
var choiceOnce sync.Once

func NewChoiceDaoInstance() *ChoiceDao {
	choiceOnce.Do(func() {
		choiceDao = &ChoiceDao{}
	})
	return choiceDao
}

func (*ChoiceDao) SelectCourse(choice *Choice) error {
	if err := db.Create(choice).Error; err != nil {
		log.Println("SelectCourse err:" + err.Error())
		return err
	}
	return nil
}
func (*ChoiceDao) IsExisted(choice *Choice) bool {
	var CHOICE Choice
	err := db.Where("student_id = ? AND course_id = ?", choice.StudentID, choice.CourseID).Find(&CHOICE).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func (*ChoiceDao) QueryCourseSelected(StudentID int) ([]int, error) {
	var choices []Choice
	var ans []int
	err := db.Where("student_id = ?", StudentID).Find(&choices).Error
	if err != nil {
		log.Println("QueryCourseSelected err:", err)
		return ans, err
	}
	ans = make([]int, len(choices))
	for i, choice := range choices {
		ans[i] = choice.CourseID
	}
	return ans, nil
}
func (*ChoiceDao) UnSelectCourse(choice *Choice) error {

	return nil
}
