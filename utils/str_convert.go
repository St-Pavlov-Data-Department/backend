package utils

import (
	"strconv"
	"strings"
)

func StrToInt64Arr(s string) (res []int64) {
	chunks := strings.Split(s, ",")

	for _, c := range chunks {
		i, _ := strconv.Atoi(c)
		res = append(res, int64(i))
	}

	return res
}

func RefToInt64Arr(ref string) (res []int64) {
	chunks := strings.Split(ref, "#")

	for _, c := range chunks {
		i, _ := strconv.Atoi(c)
		res = append(res, int64(i))
	}

	return res
}
