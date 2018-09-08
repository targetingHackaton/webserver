package utils

import "strconv"

func StrToInt(str string) int {
	if `` == str {
		return -1
	}

	value, err := strconv.Atoi(str)

	if err != nil {
		return -1
	}

	return value
}