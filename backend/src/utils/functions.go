package utils

import (
	"unicode"
)

func StrIsLetter(str string) bool {
	for _, val := range str {
		if unicode.IsLetter(val) == false {
			return false
		}
	}
	return true
}

//
func StrIsLowerLetterAndUpperLetterAndNumber(str string) bool {
	isUpper, isLower, isNumber := false, false, false
	for _, val := range str {
		if unicode.IsUpper(val) {
			isUpper = true
		} else if unicode.IsLower(val) {
			isLower = true
		} else if unicode.IsNumber(val) {
			isNumber = true
		} else {
			// 出现不是字母和数字的字符
			return false
		}
	}
	if isUpper == true && isLower == true && isNumber == true {
		return true
	} else {
		return false
	}
}
