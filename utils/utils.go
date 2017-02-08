package utils

var Id uint64 = ^uint64(0)

func GenerateId() uint64 {
	Id--
	return Id
}
