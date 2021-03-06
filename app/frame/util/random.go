package util

import (
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

func RandomUUID() string {
	u, _ := uuid.NewRandom()
	return strings.ReplaceAll(u.String(), "-", "")
}


// 获取随机字符串
func RandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}
