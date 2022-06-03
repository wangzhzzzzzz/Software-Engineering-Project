package model

import (
	"backend/src/database"
	types "backend/src/global"
	"errors"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"time"
)

type Member struct {
	UserID     int            `json:"user_id" gorm:"primary_key"`
	Nickname   string         `json:"nickname"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	UserType   types.UserType `json:"user_type"`
	CreateTime time.Time      `json:"create_time"`
	IsDeleted  bool           `json:"is_deleted"`
}

var db = database.MySqlDb

var MemberModel = Member{}

func (Member) TableName() string {
	return "member"
}
func (member *Member) CreateMember() (string, error) {
	result := Member{}
	db.Where("username = ? ", member.Username).First(&result)
	if result.UserID != 0 && !result.IsDeleted {
		log.Println(errors.New("重名错误"), "in DeleteMember")
		return "", errors.New("重名错误")
	}
	//创建一个member插入到表中
	if err := database.MySqlDb.Create(&member).Error; err != nil {
		return "", err
	}
	return strconv.Itoa(member.UserID), nil
}

func (member *Member) GetMemberByUsername(username string) (Member, error) {
	var res = Member{}
	err := database.MySqlDb.Where("username = ? ", username).First(&res).Error
	return res, err
}
func (member *Member) GetMemberByUserID(UserID string) (Member, error) {
	var res Member
	ID, _ := strconv.Atoi(UserID)
	//Scan，扫描结果至一个 struct就停止
	err := database.MySqlDb.First(&Member{}, "user_id = ? ", ID).Scan(&res).Error
	return res, err
}

// GetAllMembers 返回所有成员
func (member *Member) GetAllMembers(offset, limit int) ([]Member, error) {
	var ans []Member
	err := database.MySqlDb.Limit(limit).Offset(offset).Find(&ans).Error
	if err != nil {
		return ans, err
	}
	return ans, nil
}

func DeleteMember(user_id string) error {
	var result = Member{}
	db.Where("user_id = ? ", user_id).First(&result)
	if result.UserID == 0 {
		log.Println(errors.New("未找到该用户"), "in DeleteMember")
		return errors.New("未找到该用户")
	}
	if result.IsDeleted == true {
		log.Println(errors.New("用户已删除"), "in DeleteMember")
		return errors.New("用户已删除")
	}

	id, _ := strconv.Atoi(user_id)
	err := db.Model(&Member{}).Where("user_id = ?", id).Update("is_deleted", true).Error
	return err
}

func UpdateMember(user_id string, nickname string, username string) error {
	id, _ := strconv.Atoi(user_id)
	var result = Member{}
	db.Where(&Member{UserID: id}).First(&result)
	if result.Nickname == "" {
		return errors.New("未找到该用户")
	}
	if result.IsDeleted == true {
		return errors.New("用户已删除")
	}
	err := db.Model(&Member{}).Where("user_id = ?", user_id).Updates(Member{Nickname: nickname, Username: username}).Error
	return err
}

func AddStudentID(UserID string, rdb redis.Conn) {
	rdb.Do("SADD", "LegalStudentID", UserID)
}

func RemoveStudentID(user_id string, rdb redis.Conn) {
	rdb.Do("SREM", "LegalStudentID", user_id)
}
