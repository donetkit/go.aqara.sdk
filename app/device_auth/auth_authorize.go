package device_auth

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/donetkit/aqara.sdk/app/frame/network"

	"github.com/donetkit/aqara.sdk/app/config"
	"github.com/donetkit/aqara.sdk/app/entities"
	"github.com/donetkit/aqara.sdk/app/frame/auth_client"
	"github.com/donetkit/aqara.sdk/app/frame/util"
	"github.com/gin-gonic/gin"
)

type authAuthorize struct{}

// API管理对象
var ApiAuthAuthorize = new(authAuthorize)

// 得到预请求码 获取token
func (r *authAuthorize) OauthAuthorizeCode(c *gin.Context) {
	code := c.Query("code")
	auth_client.AuthAccessTokenClient.RegisterAccessToken(code)
}

// 设备授权接入 需要302跳转
func (r *authAuthorize) OauthAuthorize(c *gin.Context) {
	url := fmt.Sprintf("https://aiot-oauth2.aqara.cn/authorize?client_id=%s&response_type=code&redirect_uri=%s&state=%s&theme=0", config.AqaraConfig.OauthClientId, url.QueryEscape(config.AqaraConfig.AuthorizeUri), util.RandomString(16))
	c.Redirect(http.StatusFound, url)
}

// 设备授权接入 不需要跳转 body需要带上账号密码
func (r *authAuthorize) OauthAuthorize2(c *gin.Context) {
	url := fmt.Sprintf("https://aiot-oauth2.aqara.cn/authorize")
	accountRequest := &entities.AccountAuthorizeRequest{
		ClientId:     config.AqaraConfig.OauthClientId,
		ResponseType: "code",
		RedirectUri:  config.AqaraConfig.AuthorizeUri,
		State:        util.RandomString(16),
		Language:     "zh",
		Account:      "",
		Password:     "",
	}
	res, err := network.HttpPost(url, nil, accountRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
