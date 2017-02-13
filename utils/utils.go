package utils

import (
	"crypto/md5"
	"fmt"
)

var Id uint64 = ^uint64(0)

func GenerateId() string {
	Id--
	return GetHash(string(Id))
}

func GetHash(data_in string) string {
	data := []byte(data_in)
	return fmt.Sprintf("%x", md5.Sum(data))
}
