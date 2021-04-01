package entities

type DeviceQueryRequest struct {
	Did      string `json:"did"`      // 设备id
	PageNum  int    `json:"pageNum"`  // 页码，默认值1
	PageSize int    `json:"pageSize"` // 每页item个数，默认值50
}

type DeviceQueryPageRequest struct {
	PageNum  int `json:"pageNum"`  // 页码，默认值1
	PageSize int `json:"pageSize"` // 每页item个数，默认值50
}
