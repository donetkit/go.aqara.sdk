package device_container

import (
	"fmt"
	"time"

	"github.com/donetkit/aqara.sdk/app/config"
	"github.com/donetkit/aqara.sdk/app/entities"
	"github.com/donetkit/aqara.sdk/app/frame/gb"
	"github.com/donetkit/aqara.sdk/app/frame/util"
	"github.com/donetkit/aqara.sdk/app/request/auth_request"
)

type AuthAccessTokenClient struct {
	Client *AqaraClient
	Config *config.Aqara // 配置信息
}

// 初始化客户端
func NewAuthAccessTokenClient(config *config.Aqara) (client *AuthAccessTokenClient) {
	client = new(AuthAccessTokenClient)
	client.Client = new(AqaraClient)
	client.Client.Config = config
	client.Config = config
	return client
}

// 使用完整的应用凭证获取Token，如果不存在将自动注册
func (a *AuthAccessTokenClient) TryGetAccessToken() string {
	if !a.CheckRegistered() {
	}
	return a.GetAccessToken()
}

func (a *AuthAccessTokenClient) RegisterAccessToken(code string) {
	accessToken := auth_request.AccessToken(a.Config.AuthorizeUri, a.Config.OauthClientId, a.Config.OauthClientSecret, code)
	if accessToken.AccessToken != "" {
		gb.Cache.Set(fmt.Sprintf("%s:%s", gb.OauthTokenkRedisKey, a.Config.OauthClientId), accessToken, 30*24*3600)
	}
}

// 注册
func (a *AuthAccessTokenClient) Register() {
	if !a.CheckRegistered() {
	}
}

// 获取 AccessToken
func (a *AuthAccessTokenClient) GetAccessToken() string {
	return a.GetAccessTokenResult().AccessToken
}

// 获取 AccessTokenResult
func (a *AuthAccessTokenClient) GetAccessTokenResult() entities.AccessTokenResponse {
	if !a.CheckRegistered() {
		panic(fmt.Sprintf("clientId（%s）尚未注册，请先使用AccessTokenContainer.Register完成注册！", a.Config.OauthClientId))
	}
	accessToken := a.TryGetItem()
	if time.Now().After(accessToken.CreationTime.Add(time.Duration(int64(accessToken.ExpiresIn-120)) * time.Second)) {
		accessToken = auth_request.RefreshToken(config.AqaraConfig.AuthorizeUri, config.AqaraConfig.OauthClientId, config.AqaraConfig.OauthClientSecret, accessToken.RefreshToken)
		gb.Cache.Set(fmt.Sprintf("%s:%s", gb.OauthTokenkRedisKey, a.Config.OauthClientId), accessToken, 30*24*3600)
	}
	return accessToken
}

// AccessTokenBag
func (a *AuthAccessTokenClient) TryGetItem() entities.AccessTokenResponse {
	bag := entities.AccessTokenResponse{}
	obj, err1 := gb.Cache.GetString(fmt.Sprintf("%s:%s", gb.OauthTokenkRedisKey, a.Config.OauthClientId))
	if err1 != nil {
		panic(fmt.Sprintf("此appId（%s）尚未注册，请先使用AccessTokenContainer.Register完成注册！", a.Config.OauthClientId))
	}
	util.UnmarshalJson([]byte(obj), &bag)
	return bag
}

// 检查是否注册
func (a *AuthAccessTokenClient) CheckRegistered() bool {
	result, err := gb.Cache.Get(fmt.Sprintf("%s:%s", gb.OauthTokenkRedisKey, a.Config.OauthClientId))
	if err == nil && result != nil {
		return true
	}
	return false
}

// 删除缓存
func (a *AuthAccessTokenClient) RemoveFromCache() {
	gb.Cache.Del(fmt.Sprintf("%s:%s", gb.OauthTokenkRedisKey, a.Config.OauthClientId))
}
