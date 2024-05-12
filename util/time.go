package util

import (
	"strconv"
	"time"
)

const (
	// FileTime epoch is January 1, 1601
	fileTimeEpochDiff int64 = 116444736000000000
)

func UnixToFT(t time.Time) int64 {
	// Convert the time.Time value to nanoseconds since the Unix epoch
	// Default is UTC local
	nano := t.UnixNano()
	// Add the local time zone offset
	_, offset := t.Zone()
	offsetNano := int64(offset) * int64(time.Second)
	nano += offsetNano
	// Convert from nanoseconds to 100-nanosecond intervals (the unit used by FileTime)
	ft := nano / 100
	// Add the difference between the Unix and FileTime epochs
	ft += fileTimeEpochDiff
	return ft
}

func FT2Unix(ft int64) time.Time {
	// FileTime is in 100-nanosecond intervals
	// Convert to nanoseconds by multiplying by 100
	nano := int64(ft * 100)
	// FileTime epoch is January 1, 1601
	// Unix epoch is January 1, 1970
	// Calculate the difference between the two in nanoseconds
	nano -= fileTimeEpochDiff
	return time.Unix(0, nano)
}

func Time2YMDH(time time.Time) uint32 {
	timeStr := time.Format("2006010215")
	timeUInt64, _ := strconv.ParseUint(timeStr, 10, 32)
	return uint32(timeUInt64)
}
