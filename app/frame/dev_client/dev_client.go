package dev_client

import (
	"github.com/donetkit/go.aqara.sdk/app/config"
	"github.com/donetkit/go.aqara.sdk/app/device_container"
)

var DevAccessTokenClient *device_container.DevAccessTokenClient

func RegisterDevClient() *device_container.DevAccessTokenClient {
	DevAccessTokenClient = device_container.NewDevAccessTokenClient(config.AqaraConfig)
	return DevAccessTokenClient
}
