package device_token

import (
	"fmt"

	"github.com/donetkit/go.aqara.sdk/app/entities"
	"github.com/donetkit/go.aqara.sdk/app/frame/auth_client"
	"github.com/donetkit/go.aqara.sdk/app/frame/util"
)

func requestHeader() *entities.HeaderV1Request {
	head := &entities.HeaderV1Request{
		Time:        fmt.Sprintf("%d", util.FormatDateTimeUnix()),           // 	时间戳(毫秒)
		Appid:       auth_client.AuthAccessTokenClient.Config.OauthClientId, // 第三方应用的APPID
		AccessToken: auth_client.AuthAccessTokenClient.GetAccessToken(),
	}
	return head
}
