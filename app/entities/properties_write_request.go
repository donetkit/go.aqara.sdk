package entities

type PropertiesWriteDataRequest struct {
	Did  string            `json:"did"`  // 设备id
	Data map[string]string `json:"data"` // 设备属性键值对
}

type PropertiesWriteDataIntRequest struct {
	Did  string         `json:"did"`  // 设备id
	Data map[string]int `json:"data"` // 设备属性键值对
}

type PropertiesWriteRequest struct {
	Data []PropertiesWriteDataRequest `json:"data"`
}
