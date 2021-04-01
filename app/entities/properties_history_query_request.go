package entities

// 查询设备属性历史值
type PropertiesHistoryQueryDataRequest struct {
	Did string `json:"did"` // 设备ID
	//Properties []string `json:"properties"` // 	设备属性，默认查询全部
	StartTime string `json:"startTime"` // 	查询开始时间（时间戳，单位：毫秒）
	EndTime   string `json:"endTime"`   // 查询结束时间（时间戳，单位：毫秒）
	Size      int    `json:"size"`      // Max: 300, Min 10. 默认100.
}

type PropertiesHistoryQueryRequest struct {
	Data []PropertiesHistoryQueryDataRequest `json:"data"`
}

// 查询设备属性历史值
type PropertiesHistoryDateTimeDataRequest struct {
	Did string `json:"did"` // 设备ID
	//Properties []string `json:"properties"` // 	设备属性，默认查询全部
	StartTime int64 `json:"startTime"` // 	查询开始时间（时间戳，单位：毫秒）
	EndTime   int64 `json:"endTime"`   // 查询结束时间（时间戳，单位：毫秒）
	Size      int   `json:"size"`      // Max: 300, Min 10. 默认100.
}

// 查询设备属性历史值
type PropertiesHistoryDataRequest struct {
	Did        string   `json:"did"`        // 设备ID
	Properties []string `json:"properties"` // 	设备属性，默认查询全部
	Size       int      `json:"size"`       // Max: 300, Min 10. 默认100.
}
