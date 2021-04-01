package util

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// 根据签名类型，生成签名
func SignWithMd5(origin string) string {
	hashSign := Md5(origin)
	return strings.ToUpper(hex.EncodeToString(hashSign))
}

// 获取根据Key排序后的请求参数字符串
func SortSignParams(body map[string]interface{}, apiKey string) string {
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		s := fmt.Sprintf("%s=%s&", k, fmt.Sprintf("%v", body[k]))
		buffer.WriteString(s)
	}
	return buffer.String()
}
