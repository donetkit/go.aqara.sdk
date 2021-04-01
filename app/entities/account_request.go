package entities

type AccountRequest struct {
	Account  string `json:"account"`  // Aqara账号，body表单中
	Password string `json:"password"` // Aqara密码，body表单中
}

type AccountAuthorizeRequest struct {
	ClientId     string `json:"client_id"`     // 应用ID,拼接于url后
	ResponseType string `json:"response_type"` // 响应类型为code,拼接于url后
	RedirectUri  string `json:"redirect_uri"`  // 回调地址,拼接于url后
	State        string `json:"state"`         // 任意字符串，认证服务器将原样返回该参数,拼接于url后
	Language     string `json:"language"`      // 语言(en/zh),默认en,拼接于url后
	Account      string `json:"account"`       // Aqara账号，body表单中
	Password     string `json:"password"`      // Aqara密码，body表单中
}
