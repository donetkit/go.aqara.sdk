package device_sign

import (
	"fmt"

	"github.com/donetkit/aqara.sdk/app/entities"
	"github.com/donetkit/aqara.sdk/app/frame/auth_client"
	"github.com/donetkit/aqara.sdk/app/frame/util"
	"github.com/donetkit/aqara.sdk/app/request/auth_request"
)

type apiSignProperties struct{}

var ApiSignProperties = new(apiSignProperties)

// 	查询设备 /open/properties/query
func (a *apiSignProperties) PropertiesAllQuery(did string) []entities.PropertiesQueryResponse {
	response := []entities.PropertiesQueryResponse{}
	head := requestHeader()
	data := entities.PropertiesDataRequest{
		Did: did,
	}
	request := &entities.PropertiesRequest{
		Data: []entities.PropertiesDataRequest{data},
	}
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/query"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, request, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
		return response
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
			return response
		}
		util.UnmarshalJson(bytes, &response)
		return response
	}
	return nil
}

// 	查询设备 /open/properties/query     auth_device_query_response
func (a *apiSignProperties) PropertiesQuery(did, propertieName string) []entities.PropertiesQueryResponse {
	response := []entities.PropertiesQueryResponse{}
	head := requestHeader()
	data := entities.PropertiesDataQueryRequest{
		Did:        did,
		Properties: make([]string, 0),
	}
	data.Properties = append(data.Properties, propertieName)
	request := &entities.PropertiesQueryRequest{
		Data: []entities.PropertiesDataQueryRequest{data},
	}
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/query"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, request, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
		return response
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
			return response
		}
		util.UnmarshalJson(bytes, &response)
		return response
	}

	return nil

}

// 	查询设备 /open/properties/query     auth_device_query_response
func (a *apiSignProperties) PropertiesListQuery(did string, propertieName []string) []entities.PropertiesQueryResponse {
	response := []entities.PropertiesQueryResponse{}
	head := requestHeader()
	data := entities.PropertiesDataQueryRequest{
		Did:        did,
		Properties: make([]string, 0),
	}
	data.Properties = propertieName
	request := &entities.PropertiesQueryRequest{
		Data: []entities.PropertiesDataQueryRequest{data},
	}
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/query"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, request, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
		return response
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
			return response
		}
		util.UnmarshalJson(bytes, &response)
		return response
	}
	return nil

}

// 查询设备属性历史值 /open/properties/history/query
func (a *apiSignProperties) PropertiesHistoryDateTimeQuery(did string, startTime, endTime int64, size int) []entities.PropertiesQueryResponse {
	head := requestHeader()
	request := entities.PropertiesHistoryDateTimeDataRequest{
		Did: did,
		//Properties: make([]string,0),
		StartTime: startTime,
		EndTime:   endTime,
		Size:      size,
	}
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/history/query"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, request, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		response := []entities.PropertiesQueryResponse{}
		util.UnmarshalJson(bytes, &response)
		return response
	}
	return nil
}

// 查询设备属性历史值 /open/properties/history/query
func (a *apiSignProperties) PropertiesHistoryQuery(did string, propertieName []string, size int) []entities.PropertiesQueryResponse {
	head := requestHeader()
	request := entities.PropertiesHistoryDataRequest{
		Did:        did,
		Size:       size,
		Properties: make([]string, 0),
	}
	request.Properties = propertieName
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/history/query"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, request, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		response := []entities.PropertiesQueryResponse{}
		util.UnmarshalJson(bytes, &response)
		return response
	}
	return nil
}

// 提交控制指令 /open/properties/write
func (a *apiSignProperties) PropertiesWrite(did, key, value string) bool {
	head := requestHeader()
	data := entities.PropertiesWriteDataRequest{
		Did:  did,
		Data: make(map[string]string, 0),
	}
	data.Data[key] = value
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/write"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, data, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
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
			return false
		}
		response := &[]entities.PropertiesQueryResponse{}
		util.UnmarshalJson(bytes, &response)
		return true
	}
	return false
}

// 	查询设备 /open/properties/query     系统音量，取值范围为0~100
func (a *apiSignProperties) PropertiesQuerySystemVolume(did string) []entities.PropertiesQueryResponse {
	head := requestHeader()
	data := entities.PropertiesDataQueryRequest{
		Did:        did,
		Properties: make([]string, 0),
	}
	data.Properties = append(data.Properties, "system_volume")
	request := &entities.PropertiesQueryRequest{
		Data: []entities.PropertiesDataQueryRequest{data},
	}
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/query"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, request, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		response := []entities.PropertiesQueryResponse{}
		util.UnmarshalJson(bytes, &response)
		return response
	}

	return nil

}

// 	查询设备 /open/properties/query     继电器开关状态。0:关闭，1:打开,2:unknown(读)
func (a *apiSignProperties) PropertiesQueryRelayStatus(did string) []entities.PropertiesQueryResponse {
	head := requestHeader()
	data := entities.PropertiesDataQueryRequest{
		Did:        did,
		Properties: make([]string, 0),
	}
	data.Properties = append(data.Properties, "relay_status")
	request := &entities.PropertiesQueryRequest{
		Data: []entities.PropertiesDataQueryRequest{data},
	}
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/query"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, request, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		response := []entities.PropertiesQueryResponse{}
		util.UnmarshalJson(bytes, &response)
		return response
	}

	return nil

}

// 	查询设备 /open/properties/query     空调处于打开/关闭状态,0:关闭，1:打开。
func (a *apiSignProperties) PropertiesQueryOnOffStatus(did string) []entities.PropertiesQueryResponse {
	head := requestHeader()
	data := entities.PropertiesDataQueryRequest{
		Did:        did,
		Properties: make([]string, 0),
	}
	data.Properties = append(data.Properties, "relay_status")
	request := &entities.PropertiesQueryRequest{
		Data: []entities.PropertiesDataQueryRequest{data},
	}
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/query"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, request, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := &entities.BaseResponse{}
	util.UnmarshalJson(bytes, resp)
	if resp.Code == 0 {
		bytes, err = util.Decryptbase64DecodeString(resp.Result.(string), []byte(auth_client.AuthAccessTokenClient.Config.DevClientSecret))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		response := []entities.PropertiesQueryResponse{}
		util.UnmarshalJson(bytes, &response)
		return response
	}

	return nil

}

// 提交控制指令 /open/properties/write 系统音量，取值范围为0~100
func (a *apiSignProperties) PropertiesWriteSystemVolume(did string, value int) bool {
	head := requestHeader()
	data := entities.PropertiesWriteDataRequest{
		Did:  did,
		Data: make(map[string]string, 0),
	}
	data.Data["system_volume"] = fmt.Sprintf("%d", value)
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/write"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, data, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
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
			return false
		}
		response := &[]entities.PropertiesQueryResponse{}
		util.UnmarshalJson(bytes, &response)
		return true
	}
	return false
}

// 提交控制指令 /open/properties/write  继电器开关状态。0:关闭，1:打开,2:unknown(读)
func (a *apiSignProperties) PropertiesWriteRelayStatus(did string, value int) bool {
	head := requestHeader()
	data := entities.PropertiesWriteDataRequest{
		Did:  did,
		Data: make(map[string]string, 0),
	}
	data.Data["relay_status"] = string(value)
	url := "https://aiot-open-3rd.aqara.cn/v2.0/open/properties/write"
	bytes, err := auth_request.DoAiotHttpPostSignHeader(url, head, data, auth_client.AuthAccessTokenClient.Config.DevClientSecret)
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
			return false
		}
		response := &[]entities.PropertiesQueryResponse{}
		util.UnmarshalJson(bytes, &response)
		return true
	}
	return false
}
