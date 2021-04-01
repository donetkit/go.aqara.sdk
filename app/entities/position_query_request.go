package entities


// 查询位置
type PositionQueryRequest struct {
	PositionId string `json:"position_id"`   // 位置ID，为空时查询该用户下所有位置
	PageNum int `json:"page_num"`   // 页码，默认值1
	PageSize int `json:"page_size"` // 每页item个数，默认值50
}
