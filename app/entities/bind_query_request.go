package entities

// 设备入网状态查询
type BindQueryRequest struct {
	BindKey string `json:"bindkey"`  // 入网凭证获取，有效时间10分钟（16位字符串[0-9a-zA-Z]）
	Did string `json:"did"`  // 设备id
}
