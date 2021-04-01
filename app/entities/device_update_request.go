package entities

type DeviceUpdateRequest struct {
	Did string `json:"did"`  // 设备id
	Name string `json:"name"` // 	设备名称
}
