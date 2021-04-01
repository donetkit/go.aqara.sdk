package entities

// 请求授权码
type AqaraAuthorizeRequest struct {
	ClientId     string `json:"client_id"`     //第三方应用ID，AppID
	ResponseType string `json:"response_type"` // 返回类型，按照OAuth 2.0 标准，取值为code
	RedirectUri  string `json:"redirect_uri"`  //第三方应用注册的重定向URI
	State        string `json:"state"`         // 取值为任意字符串，认证服务器将原样返回该参数
	Language string `json:"language"` // 语言(en/zh),默认en,拼接于url后
	Account string `json:"account"` // Aqara账号，body表单中
	Password string `json:"password"` // Aqara密码，body表单中
}
