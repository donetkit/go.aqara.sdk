package entities

// 查询设备当前属性值
type PropertiesQueryResponse struct {
	Property string `json:"property"` // 设备属性
	Time string `json:"time"` // 设备属性变更时间
	Value string `json:"value"` // 设备属性值
	Did string `json:"did"`  // 设备id
}
