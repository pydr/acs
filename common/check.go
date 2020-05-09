package common

import (
	"regexp"
)

func CheckUsername(username string) bool {
	pattern := `^[a-zA-Z0-9_]{3,18}$`
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(username)
}

func CheckPwd(password string) bool {
	pattern := `^[a-zA-Z0-9_]{8,16}$`
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(password)
}

func CheckMobile(mobile string) bool {
	pattern := `^1([38][0-9]|4[579]|5[0-3,5-9]|6[6]|7[0135678]|9[89])\d{8}$`
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(mobile)
}
