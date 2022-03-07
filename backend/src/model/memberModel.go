package model

import (
	"backend/src/database"
	types "backend/src/global"
)

type Member struct {
	UserID    int            `json:"user_id" gorm:"primary_key"`
	Nickname  string         `json:"nickname"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	UserType  types.UserType `json:"user_type"`
	IsDeleted bool           `json:"is_deleted"`
}

func (Member) TableName() string {
	return "member"
}
func (member *Member) GetMemberByUsername(username string) (Member, error) {
	var res = Member{}
	err := database.MySqlDb.Where("username = ? ", username).First(&res).Error
	return res, err
}
