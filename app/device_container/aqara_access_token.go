package device_container

// 授权
type AqaraAccessToken interface {
	// 获取授权码
	GetAccessToken()
}
