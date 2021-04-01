package entities

// 获取访问令牌
type AccessTokenResponse struct {
	BaseRequest
	AccessToken  string `json:"access_token"`  // 访问令牌，第三方应用访问AIOT开放服务的凭证
	ExpiresIn    int    `json:"expires_in"`    // 访问令牌的剩余有效时间，单位为秒
	TokenType    string `json:"token_type"`    // 根据OAuth 2.0 标准，取值为bearer
	OpenId       string `json:"openId"`       // 授权用户的唯一标识
	RefreshToken string `json:"refresh_token"` // 刷新令牌，用于刷新访问令牌，有效期为30天
	State        string `json:"state"`         // 取值为任意字符串，认证服务器将原样返回该参数

}
