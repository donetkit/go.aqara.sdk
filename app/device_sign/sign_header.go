package device_sign

import (
	"fmt"

	"github.com/donetkit/aqara.sdk/app/entities"
	"github.com/donetkit/aqara.sdk/app/frame/auth_client"
	"github.com/donetkit/aqara.sdk/app/frame/util"
)

func requestHeader() *entities.HeaderRequest {
	head := &entities.HeaderRequest{
		Time:  fmt.Sprintf("%d", util.FormatDateTimeUnix()), // 	时间戳(毫秒)
		Appid: auth_client.AuthAccessTokenClient.Config.DevClientId,
		Lang:  "zh",
	}
	return head
}
