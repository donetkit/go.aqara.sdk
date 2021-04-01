package entities

type AuthDeviceQueryResponse struct {
	Code      int          `json:"code"`
	RequestId string       `json:"requestId"`
	Message   string       `json:"message"`
	Result    DataResponse `json:"result"`
}

type DataResponse struct {
	Data []AuthDeviceQueryDataResponse `json:"data"`
}

type AuthDeviceQueryDataResponse struct {
	Did string `json:"did"` // 设备id
	Name string `json:"name"` //设备名称
	Model string `json:"model"` // 设备类型
	ParentId string `json:"parentId"` // 父设备（网关）ID，若是网关则该字段为空
	PositionId string `json:"positionId"` // 设备位置
	State int `json:"state"` // 是否在线（0-离线，1-在线）
	RegisterTime string `json:"registerTime"` // 首次注册时间
}