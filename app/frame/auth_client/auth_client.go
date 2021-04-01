package auth_client

import (
	"github.com/donetkit/go.aqara.sdk/app/config"
	"github.com/donetkit/go.aqara.sdk/app/device_container"
)

var AuthAccessTokenClient *device_container.AuthAccessTokenClient

func RegisterAuthClient() *device_container.AuthAccessTokenClient {
	AuthAccessTokenClient = device_container.NewAuthAccessTokenClient(config.AqaraConfig)
	return AuthAccessTokenClient
}
