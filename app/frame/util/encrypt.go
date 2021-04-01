package util

import (
	"crypto/md5"
)

// MD5
func Md5(str string) []byte {
	hash := md5.New()
	_, _ = hash.Write([]byte(str))
	return hash.Sum(nil)
}
