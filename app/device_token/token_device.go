package device_token

import (
	"fmt"

	"github.com/donetkit/go.aqara.sdk/app/entities"
	"github.com/donetkit/go.aqara.sdk/app/frame/auth_client"
	"github.com/donetkit/go.aqara.sdk/app/frame/util"
	"github.com/donetkit/go.aqara.sdk/app/request/auth_request"
)

type apiDevice struct{}

var ApiDevice = new(apiDevice)

// 	查询设备 /open/device/query
func (a *apiDevice) DeviceQueryDid(did string, pageNum, pageSize int) []entities.AuthDeviceQueryDataResponse {
	head := requestHeader()
	body := &entities.DeviceQueryRequest{
		Did:      did,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/query"
	bytes, err := auth_request.DoAiotHttpGetHeader(url, head, body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	response := entities.AuthDeviceQueryResponse{}
	util.UnmarshalJson(bytes, &response)
	if response.Code == 0 {
		return response.Result.Data
	}
	return nil
}

// 	查询设备 /open/device/query
func (a *apiDevice) DeviceQueryPage(pageNum, pageSize int) []entities.AuthDeviceQueryDataResponse {
	head := requestHeader()
	body := &entities.DeviceQueryPageRequest{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/query"
	bytes, err := auth_request.DoAiotHttpGetHeader(url, head, body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	response := entities.AuthDeviceQueryResponse{}
	util.UnmarshalJson(bytes, &response)
	if response.Code == 0 {
		return response.Result.Data
	}
	return nil
}

// 	查询设备 /open/device/query
func (a *apiDevice) DeviceQuery() []entities.AuthDeviceQueryDataResponse {
	head := requestHeader()
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/query"
	bytes, err := auth_request.DoAiotHttpGetHeader(url, head, nil)
	if err != nil {
		fmt.Println(err)
	}
	response := entities.AuthDeviceQueryResponse{}
	util.UnmarshalJson(bytes, &response)
	if response.Code == 0 {
		return response.Result.Data
	}
	return nil
}

// 	更新设备信息 /open/device/update
func (a *apiDevice) UpdateDeviceName(did, name string) bool {
	head := requestHeader()
	body := &entities.DeviceUpdateRequest{
		Did:  did,  // 		设备id
		Name: name, // 		设备名称
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/update"
	bytes, err := auth_request.DoAiotHttpPostHeader(url, head, body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
		}
		return true
	}
	return false
}

// 	网关开启子设备入网 	/open/device/connect
func (a *apiDevice) DeviceConnect(did string) bool {
	head := requestHeader()
	body := &entities.DeviceConnectRequest{
		Did: did, // 		设备id
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/connect"
	bytes, err := auth_request.DoAiotHttpPostHeader(url, head, body)
	if err != nil {
		fmt.Println(err)
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		return true
	}
	return false

}

// 	网关关闭子设备入网 	/open/device/connect/stop
func (a *apiDevice) DeviceConnectStop(did string) bool {
	head := requestHeader()
	body := &entities.DeviceConnectRequest{
		Did: did, // 		设备id
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/connect/stop"
	bytes, err := auth_request.DoAiotHttpPostHeader(url, head, body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		return true
	}
	return false

}

// 	查询网关下子设备信息 	/open/device/child/query
func (a *apiDevice) DeviceChildQuery(did string) {
	head := requestHeader()
	body := &entities.DeviceConnectRequest{
		Did: did, // 		设备id
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/child/query"
	bytes, err := auth_request.DoAiotHttpPostHeader(url, head, body)
	if err != nil {
		fmt.Println(err)
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)

}

// 	设备解绑(网关和子设备) 	/open/device/unbind
func (a *apiDevice) DeviceUnBind(did string) {
	head := requestHeader()
	body := &entities.DeviceUnbindRequest{
		Did:    did, // 		设备id
		Option: 1,   // 0-保留自动化场景信息，1-清除自动化场景信息
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/unbind"
	_, err := auth_request.DoAiotHttpPostHeader(url, head, body)
	if err != nil {
		fmt.Println(err)
	}

}

// 	查询资源	/open/resource/query
func (a *apiDevice) DeviceResourceQuery(did string, attrs []string) {
	head := requestHeader()
	body := &entities.DeviceResourceRequest{
		Did:   did,        // 		设备id
		Attrs: []string{}, // 0-保留自动化场景信息，1-清除自动化场景信息
	}
	body.Attrs = attrs
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/device/unbind"
	_, err := auth_request.DoAiotHttpPostHeader(url, head, body)
	if err != nil {
		fmt.Println(err)
	}

}
