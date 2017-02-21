package utils

import "fmt"

var userCount uint64 = 0
var postCount uint64 = 0

func GenerateId(isUser bool) string {
	if isUser {
		userCount++
		return fmt.Sprintf("%d", userCount)
	} else {
		postCount++
		return fmt.Sprintf("%d", postCount)
	}
}
