package utils

import (
	"strings"
)

func Rightful(data string, min, max int) bool {
	data = strings.Replace(data, " ", "", -1)
	//判断pwd
	if max == 0 {
		if len(data) < min {
			return false
		}
	} else {
		if len(data) < min || len(data) > max {
			return false
		}
	}
	return true
}
