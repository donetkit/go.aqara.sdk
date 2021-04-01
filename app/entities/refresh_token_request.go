package entities

// 刷新访问令牌
type RefreshTokenRequest struct {
	ClientId     string `json:"client_id"`     // 第三方应用ID，AppID
	ClientSecret string `json:"client_secret"` // 第三方应用秘钥，AppKey
	GrantType    string `json:"grant_type"`    // 根据OAuth 2.0 标准，取值为refresh_token
	RefreshToken string `json:"refresh_token"` // 上一步请求获得的授权码
	RedirectUri  string `json:"redirect_uri"`  // 上一步请求中设置的redirect_uri参数
	State string `json:"state"` // 任意字符串，认证服务器将原样返回该参数
}
