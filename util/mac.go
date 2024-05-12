package util

import (
	"fmt"
	"strconv"
	"strings"
)

func MAC2Bytes(mac string) ([]byte, error) {
	strSlice := strings.Split(mac, "-")
	bytes := make([]byte, 6)
	for i, s := range strSlice {
		val, err := strconv.ParseUint(s, 16, 8)
		if err != nil {
			return nil, err
		}
		bytes[i] = byte(val)
	}
	return bytes, nil
}

func Bytes2MAC(bytes []byte) string {
	var build strings.Builder
	for i, b := range bytes {
		build.WriteString(fmt.Sprintf("%02x", b))
		if i != 5 {
			build.WriteString("-")
		}
	}
	return build.String()
}
