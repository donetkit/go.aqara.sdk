package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/donetkit/aqara.sdk/app/frame/util"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 5 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
		},
	}
}

// 发送GET请求
func HttpGet(url string) ([]byte, error) {
	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}

// 发送POST请求(JSON)
func HttpPostJson(url string, body interface{}) ([]byte, error) {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return httpPost(url, "application/json", string(bodyStr))
}

// 发送通用的POST请求
func HttpPostFormMap(url string, body interface{}) ([]byte, error) {
	jsonStr := util.MarshalJson(body)
	val := util.JSONToMapString(jsonStr)
	return httpPostForm(url, val)
}

// 发送通用的POST请求
func httpPost(url string, contentType string, body string) ([]byte, error) {
	rsp, err := client.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}

// 发送通用的POST请求
func httpPostForm(httpUrl string, body map[string]string) ([]byte, error) {
	dataVal := url.Values{}
	for key, val := range body {
		dataVal.Add(key, val)
	}
	rsp, err := client.PostForm(httpUrl, dataVal)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)

}

// 发送带Header的Get请求
func HttpGetHeaderMap(httpUrl string, header map[string]interface{}, body interface{}) ([]byte, error) {
	clientGet := &http.Client{}
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest("GET", httpUrl, nil)
	} else {
		httpUrl = fmt.Sprintf("%s?%s", httpUrl, util.StructToUrlParams(body))
		req, err = http.NewRequest("GET", httpUrl, nil)
	}
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	//增加header选项
	for key, val := range header {
		req.Header.Add(key, val.(string))
	}
	resp, err := clientGet.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// 发送通用的POST请求
func HttpPostHeaderMap(httpUrl string, header map[string]interface{}, body interface{}) ([]byte, error) {
	clientGet := &http.Client{}
	jsonStr := util.MarshalJson(body)
	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	//增加header选项
	for key, val := range header {
		req.Header.Add(key, val.(string))
	}
	resp, err := clientGet.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// 发送通用的POST请求
func HttpPostHeaderMapPayload(httpUrl string, header map[string]interface{}, body interface{}) ([]byte, error) {
	clientGet := &http.Client{}
	jsonStr := header["Payload"].(string)
	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	//增加header选项
	for key, val := range header {
		req.Header.Add(key, val.(string))
	}
	resp, err := clientGet.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// 发送通用的POST请求
func HttpPost(httpUrl string, header map[string]interface{}, body interface{}) (*http.Response, error) {
	clientGet := &http.Client{}
	valMaps := util.JSONToMapString(util.MarshalJson(body))
	dataVal := url.Values{}
	for key, val := range valMaps {
		dataVal.Add(key, val)
	}
	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(dataVal.Encode()))
	if err != nil {
		panic(err)
	}
	//增加header选项
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if header != nil {
		for key, val := range header {
			req.Header.Add(key, val.(string))
		}
	}
	return clientGet.Do(req)
}
