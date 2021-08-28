package util

import (
	"regexp"
	"strings"
)

func CheckNameFormat(username string) (bool, string) {
	length := strings.Count(username, "") - 1
	if length < 5 {
		return false, "Username is too short."
	} else if length > 16 {
		return false, "Username is too long."
	}
	match, _ := regexp.MatchString(`^[a-zA-Z][0-9a-zA-Z_-]{4,15}$`, username)
	if !match {
		return match, "用户名只能包含_-,首位必须为字母。"
	}
	return match, ""
}
