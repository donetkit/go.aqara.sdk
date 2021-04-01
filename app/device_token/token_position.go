package device_token

import (
	"fmt"

	"github.com/donetkit/aqara.sdk/app/entities"
	"github.com/donetkit/aqara.sdk/app/frame/auth_client"
	"github.com/donetkit/aqara.sdk/app/frame/util"
	"github.com/donetkit/aqara.sdk/app/request/auth_request"
)

type apiPosition struct{}

var ApiPosition = new(apiPosition)

// 查询位置 /open/position/query
func (a *apiPosition) PositionQuery() {
	head := requestHeader()
	body := &entities.PositionQueryRequest{
		PositionId: "",
		PageNum:    1,
		PageSize:   10,
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/position/query"
	bytes, err := auth_request.DoAiotHttpGetHeader(url, head, body)
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
}

// 创建位置 /open/position/add
func (a *apiPosition) PositionAdd() {
	head := requestHeader()
	body := &entities.PositionAddRequest{
		PositionName:     "位置名称",
		Description:      "位置描述",
		ParentPositionId: "",
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/position/add"
	bytes, err := auth_request.DoAiotHttpPostHeader(url, head, body)
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
		fmt.Println("请求结果：" + string(bytes))
	}

}

// 删除位置 	/open/position/delete
func (a *apiPosition) PositionDelete() {
	head := requestHeader()
	body := &entities.PositionDeleteRequest{
		PositionId: "",
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/position/delete"
	bytes, err := auth_request.DoAiotHttpPostHeader(url, head, body)
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
		fmt.Println("请求结果：" + string(bytes))
	}
}

// 	更新位置 	/open/position/update
func (a *apiPosition) PositionUpdate() {
	head := requestHeader()
	body := &entities.PositionUpdateRequest{
		PositionId:   "",
		PositionName: "",
		Description:  "",
	}
	url := "https://aiot-open-3rd.aqara.cn/3rd/v1.0/open/position/update"
	bytes, err := auth_request.DoAiotHttpPostHeader(url, head, body)
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
		fmt.Println("请求结果：" + string(bytes))
	}
}
