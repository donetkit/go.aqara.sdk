package dev_request

import (
	"encoding/json"
	"fmt"
	"github.com/donetkit/aqara.sdk/app/frame/network"
	"github.com/donetkit/aqara.sdk/app/frame/util"
	"strings"
)

// 构建Body
func buildBody(bodyObj interface{}, appkey string) (body map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, _ := json.Marshal(bodyObj)
	body = make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)
	// 生成签名
	var sign string
	sign = localSign(body, appkey)
	body["Content-Type"] = "application/json"
	body["sign"] = sign
	return
}

// 构建Body
func buildHeaders(headerObj interface{}, appkey string) (heads map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	headerJson, _ := json.Marshal(headerObj)
	heads = make(map[string]interface{})
	_ = json.Unmarshal(headerJson, &heads)
	// 生成签名
	var sign string
	sign = localSign(heads, appkey)
	heads["Sign"] = strings.ToLower(sign)
	return
}

// 本地通过支付参数计算签名值
func localSign(body map[string]interface{}, appkey string) string {
	signStr := util.SortSignParams(body,appkey)
	signStr = strings.ToLower(signStr)
	signStr= fmt.Sprintf("%s%s",signStr,appkey)
	return util.SignWithMd5(signStr)
}

// 向aiot发送请求
func DoAiotHttp(url string, bodyObj interface{}, appkey string) (bytes []byte, err error) {
	// 转换参数
	body, err := buildBody(bodyObj,appkey)
	if err != nil {
		return
	}
	// 发起请求
	bytes, err = network.HttpPostJson(url, body)
	return
}


// 向aiot发送请求
func DoAiotHttpGet(url string) (bytes []byte, err error) {
	// 发起请求
	bytes, err = network.HttpGet(url)
	return
}


// 向aiot发送请求
func DoAiotHttpGetHeader(url string, header interface{}, appkey string) (bytes []byte, err error) {
	headers,err := buildHeaders(header,appkey)
	if err !=nil {
		fmt.Println(err.Error())
	}
	// 发起请求
	bytes, err = network.HttpGetHeaderMap(url, headers,nil)
	return
}
