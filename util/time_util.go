package util

import "time"

// Time2String 时间转字符串
func Time2String(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
