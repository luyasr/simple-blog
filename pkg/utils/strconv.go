package utils

import (
	"strconv"
)

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Int64ToString(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}
