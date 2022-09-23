package utils

import "strconv"

func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}
