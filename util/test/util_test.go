package test

import (
	"encoding/binary"
	"fmt"
	"goms/util"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	s := ""
	bytes := util.Str2Bytes(s)
	fmt.Println(bytes)
}

func TestUint32(t *testing.T) {
	buf := []byte{0, 0, 0, 0}
	val := binary.LittleEndian.Uint32(buf)
	fmt.Println(val)
}

func TestUint32ToBytes(t *testing.T) {
	var number uint32 = 1980835063
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, number)
	fmt.Println(buf)
}

func TestFT2Unix(t *testing.T) {
	var ft int64 = 512796963281250
	unix := util.FT2Unix(ft)
	fmt.Println(unix)
}

func TestYMDH2Time(t *testing.T) {
	buf := []byte{145, 143, 43, 120}
	number := binary.LittleEndian.Uint32(buf)
	date := util.YMDH2Time(number)
	fmt.Println(date)
}
