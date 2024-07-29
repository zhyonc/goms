package util

import (
	"strconv"
	"strings"
)

func FormatData(data []byte) string {
	length := len(data)
	if length == 0 {
		return ""
	}
	lastIndex := length - 1
	var builder strings.Builder
	for i, v := range data {
		builder.WriteString(strconv.Itoa(int(v)))
		if i < lastIndex {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

func Str2Bytes(s string) string {
	str := strings.Split(s, " ")
	strs := make([]string, len(str))
	for i := 0; i < len(strs); i++ {
		hex, _ := strconv.ParseUint(str[i], 16, 8)
		strs[i] = strconv.Itoa(int(hex))
	}
	return strings.Join(strs, ",")
}
