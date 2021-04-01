package device_sign

import (
	"fmt"

	"github.com/donetkit/aqara.sdk/app/entities"
	"github.com/donetkit/aqara.sdk/app/frame/auth_client"
	"github.com/donetkit/aqara.sdk/app/frame/util"
	"github.com/donetkit/aqara.sdk/app/request/auth_request"
)

type apiSignDevice struct{}

var ApiSignDevice = new(apiSignDevice)

// 	查询设备 /open/dev/bind/get
func (a *apiSignDevice) DeviceBindGet() {
	head := requestHeader()
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/dev/bind/get"
	_, err := auth_request.DoAiotHttpGetSignHeader(url, head, nil, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
	}
}

// 设备入网状态查询 /open/dev/bind/query
func (a *apiSignDevice) DeviceBindquery(bindKey, did string) {
	head := requestHeader()
	body := &entities.BindQueryRequest{
		BindKey: bindKey,
		Did:     did,
	}
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/dev/bind/query"

	bytes, err := auth_request.DoAiotHttpGetSignHeader(url, head, body, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("请求结果：" + string(bytes))
}
