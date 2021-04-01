package entities

type DeviceUnbindRequest struct {
	Did    string `json:"did"`    // 设备id
	Option int    `json:"option"` // 0-保留自动化场景信息，1-清除自动化场景信息
}

type DeviceResourceRequest struct {
	Did   string   `json:"did"`   // 设备id
	Attrs []string `json:"attrs"` // 资源别名
}
