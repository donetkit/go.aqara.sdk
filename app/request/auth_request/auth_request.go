package auth_request

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/donetkit/go.aqara.sdk/app/frame/network"
	"github.com/donetkit/go.aqara.sdk/app/frame/util"
)

type buildOauthHandler func(bodyObj interface{}) (body map[string]interface{}, err error)

// 构建Body
func buildBody(bodyObj interface{}, appKey string) (body map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, _ := json.Marshal(bodyObj)
	body = make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)
	// 生成签名
	var sign string
	sign = localSign(body, appKey)
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

// 构建Body
func buildHeadersNoSign(headerObj interface{}) (heads map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	headerJson, _ := json.Marshal(headerObj)
	heads = make(map[string]interface{})
	_ = json.Unmarshal(headerJson, &heads)
	return
}

// 本地通过支付参数计算签名值
func localSign(body map[string]interface{}, appkey string) string {
	signStr := util.SortSignParams(body, appkey)
	signStr = strings.ToLower(signStr)
	signStr = fmt.Sprintf("%s%s", signStr, appkey)
	return util.SignWithMd5(signStr)
}

// 向aiot发送请求
func DoAiotHttpPostJson(url string, bodyObj interface{}, appKey string) (bytes []byte, err error) {
	// 转换参数
	body, err := buildBody(bodyObj, appKey)
	if err != nil {
		panic(err)
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

// 向aiot发送请求 HttpGetHeaderBodyMap
func DoAiotHttpGetSignHeader(url string, header interface{}, body interface{}, appKey string) (bytes []byte, err error) {
	headers, err := buildGetHeaders(header, appKey)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	bytes, err = network.HttpGetHeaderMap(url, headers, body)
	return
}

// 向aiot发送请求 HttpGetHeaderBodyMap
func DoAiotHttpPostSignHeader(url string, header interface{}, body interface{}, appKey string) (bytes []byte, err error) {
	headers, err := buildGetHeadersBody(header, appKey, body)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	bytes, err = network.HttpPostHeaderMapPayload(url, headers, body)
	return
}

// 构建Body
func buildGetHeadersBody(headerObj interface{}, appkey string, body interface{}) (heads map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	jsonBody := util.MarshalJson(body)
	payload, err := util.Encrypt([]byte(jsonBody), util.GetAesIv([]byte(appkey)))
	headerJson, _ := json.Marshal(headerObj)
	heads = make(map[string]interface{})
	heads["Payload"] = base64.StdEncoding.EncodeToString(payload)
	_ = json.Unmarshal(headerJson, &heads)
	sign := localSign(heads, appkey) // 生成签名
	heads["Sign"] = strings.ToLower(sign)
	return
}

// 构建Body
func buildGetHeaders(headerObj interface{}, appkey string) (heads map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	headerJson, _ := json.Marshal(headerObj)
	heads = make(map[string]interface{})
	_ = json.Unmarshal(headerJson, &heads)
	sign := localSign(heads, appkey) // 生成签名
	heads["Sign"] = sign
	return
}

// 向aiot发送请求 HttpGetHeaderBodyMap
func DoAiotHttpGetHeader(url string, header interface{}, body interface{}) (bytes []byte, err error) {
	headers, err := buildHeadersNoSign(header)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	bytes, err = network.HttpGetHeaderMap(url, headers, body)
	return
}

// 向aiot发送请求
func DoAiotHttpPostHeader(url string, header interface{}, body interface{}) (bytes []byte, err error) {
	headers, err := buildHeadersNoSign(header)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// 发起请求
	bytes, err = network.HttpPostHeaderMap(url, headers, body)
	return
}
