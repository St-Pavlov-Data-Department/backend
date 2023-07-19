package utils

import (
	"strconv"
	"strings"
)

func StoInt64Arr(s string) (res []int64) {
	chunks := strings.Split(s, ",")

	for _, c := range chunks {
		i, _ := strconv.Atoi(c)
		res = append(res, int64(i))
	}

	return res
}
