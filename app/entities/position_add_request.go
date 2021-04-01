package entities

// 创建位置
type PositionAddRequest struct {
	PositionName string `json:"position_name"` // 位置名称
	Description string `json:"description"` // 位置描述
	ParentPositionId string `json:"parent_position_id"` // 父位置id，创建顶级位置时为空
}
