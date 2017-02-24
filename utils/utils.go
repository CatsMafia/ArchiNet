package utils

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"strings"
)

var kekCount uint64 = 0

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
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
