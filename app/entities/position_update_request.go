package entities

// 更新位置
type PositionUpdateRequest struct {
	PositionId string `json:"position_id"` //位置id
	PositionName string `json:"position_name"` //	位置名称
	Description string `json:"description"` //位置描述
}
