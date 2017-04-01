package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

var userCount uint64 = 0
var lolsCount uint64 = 0

func GenerateUserId() string {
	userCount++
	return strconv.FormatUint(userCount, 10)
}

func GenerateId() string {
	lolsCount++
	return fmt.Sprintf("%d", lolsCount)
}
func FindSubStr(str string, start, end string) string {
	if strings.ContainsAny(str, string(start)+" & "+string(end)) {
		str += " "
		flag := false
		var temp int
		res := ""
		for i, c := range str {
			if string(c) == start {
				if flag {
					res += str[temp:i]
					temp = i
					flag = true
				} else {
					temp = i
					flag = true
				}
			} else if string(c) == end {
				if flag {
					res += str[temp:i]
					flag = false
				}
			}
		}
		return res
	} else {
		return ""
	}
}

func IsIn(sl []string, str string) bool {
	for _, h := range sl {
		if h == str {
			return true
		}
	}
	return false
}

func GetHash(in string) string {
	data := []byte(in)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func RemoveElemString(list []string, item string) []string {
	for i, v := range list {
		if v == item {
			copy(list[i:], list[i+1:])
			list = list[:len(list)-1]
		}
	}
	return list
}
