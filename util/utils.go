/**
 * @Author: xianxiong
 * @Date: 2020/11/1 14:39
 */

package util

import (
	"strconv"
	"time"
)

//StrToInt string 转int
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

//StrToUInt string 转uint
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

// StrToTime 字符串转time
func StrToTime(s string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(timeLayout, s, time.Local)
	return t
}

// StrToDate 字符串转time
func StrToDate(s string) time.Time {
	timeLayout := "2006-01-02"
	t, _ := time.ParseInLocation(timeLayout, s, time.Local)
	return t
}

//StrToTimePtr ..
func StrToTimePtr(str string) *time.Time {
	t, e := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if e != nil {
		return nil
	}
	return &t
}
