package util

import (
	"time"
)

// 格式化时间，时间戳（毫秒）
func FormatDateTimeUnix() int64 {
	return time.Now().UnixNano() / 1000000 //  strconv.FormatInt(, 10)
}
