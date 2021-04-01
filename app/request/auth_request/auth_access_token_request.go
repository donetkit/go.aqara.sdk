package auth_request

import (
	"time"

	"github.com/donetkit/aqara.sdk/app/entities"
	"github.com/donetkit/aqara.sdk/app/frame/network"
	"github.com/donetkit/aqara.sdk/app/frame/util"
)

// 获取访问令牌
func AccessToken(authorizeUri, clientId, clientSecret, code string) entities.AccessTokenResponse {
	accessTokenRequest := &entities.AccessTokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "authorization_code",
		Code:         code,
		RedirectUri:  authorizeUri,
		State:        util.RandomString(16),
	}
	url := "https://aiot-oauth2.aqara.cn/access_token"
	bytes, err := network.HttpPostFormMap(url, accessTokenRequest)
	if err != nil {
		panic(err)
	}
	accessToken := entities.AccessTokenResponse{}
	accessToken.CreationTime = time.Now()
	util.UnmarshalJson(bytes, &accessToken)
	return accessToken
}

// 刷新访问令牌 // https://aiot-oauth2.aqara.cn/access_token
func RefreshToken(redirectUri string, clientId string, clientSecret string, refreshToken string) entities.AccessTokenResponse {
	refreshTokenRequest := &entities.RefreshTokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
		RedirectUri:  redirectUri,
		State:        util.RandomString(16),
	}
	url := "https://aiot-oauth2.aqara.cn/access_token"
	bytes, err := network.HttpPostFormMap(url, refreshTokenRequest)
	if err != nil {
		panic(err)
	}
	accessToken := entities.AccessTokenResponse{}
	accessToken.CreationTime = time.Now()
	util.UnmarshalJson(bytes, &accessToken)
	return accessToken
}
