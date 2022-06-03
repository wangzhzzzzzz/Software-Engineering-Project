package validate

import global "backend/src/global"

// IsValidCreateMember 验证传入的CreateMemberRequest参数是否合法
func IsValidCreateMember(param *global.CreateMemberRequest) bool {
	if len(param.Nickname) < 4 || len(param.Nickname) > 20 {
		return false
	}
	if len(param.Username) < 8 || len(param.Username) > 20 {
		return false
	}
	if len(param.Password) < 8 || len(param.Password) > 20 {
		return false
	}
	if param.UserType < 1 || param.UserType > 3 {
		return false
	}
	return true
}

//
