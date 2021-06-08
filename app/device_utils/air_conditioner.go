package device_utils

import (
	"fmt"

	"github.com/donetkit/go.aqara.sdk/app/device_utils/common"
)

// 开关
func airConditionerSwitch(data string) string {
	result := "reserve"
	res := common.BinHex(data)
	switch res {
	case "0":
		result = "关"
		break
	case "1":
		result = "开"
		break
	case "2":
		result = "toggle"
		break
	case "E":
		result = "circle"
		break
	case "F":
		result = "invalid"
		break
	}
	return result
}

// 模式
func airConditionerModel(data string) string {
	result := "reserve" // 0: heat; 1: cool; 2: auto; 3: dry; 4: wind; E: circle; F: invalid; else: reserve	模式
	res := common.BinHex(data)
	switch res {
	case "0":
		result = "制热"
		break
	case "1":
		result = "制冷"
		break
	case "2":
		result = "自动"
		break
	case "3":
		result = "除湿"
		break
	case "4":
		result = "送风"
		break
	case "E":
		result = "circle"
		break
	case "F":
		result = "invalid"
		break
	}
	return result
}

// 风速
func airConditionerWindSpeed(data string) string {
	result := "reserve" // 0: low; 1: middle; 2: high; 3: auto; E: circle; F: invalid; else: reserve	风速
	res := common.BinHex(data)
	switch res {
	case "0":
		result = "低"
		break
	case "1":
		result = "中"
		break
	case "2":
		result = "高"
		break
	case "3":
		result = "自动"
		break
	case "E":
		result = "circle"
		break
	case "F":
		result = "invalid"
		break
	}
	return result
}

// 风向
func airConditionerWindDirection(data string) string {
	result := "reserve" // 0: horizontal; 1: vertical; 2: circle; 3: invalid;	风向
	res := common.BinHex(data)
	switch res {
	case "0":
		result = "左右"
		break
	case "1":
		result = "上下"
		break
	case "2":
		result = "circle"
		break
	case "3":
		result = "invalid"
		break
	}
	return result
}

// 扫风
func airConditionerSweeping(data string) string {
	result := "reserve" // 0: swing; 1: fix; 2: circle; 3: invalid;	扫风
	res := common.BinHex(data)
	switch res {
	case "0":
		result = "摆动"
		break
	case "1":
		result = "固定"
		break
	case "2":
		result = "circle"
		break
	case "3":
		result = "invalid"
		break
	}
	return result
}

// 温度
func airConditionerTemperature(data string) string {
	result := fmt.Sprintf("%d", common.BinDec(data)) // 0 ~ 240; 243: up; 244: down; FF: invalid	温度
	res := common.BinHex(data)                       //
	switch res {
	case "243":
		result = "up"
		break
	case "244":
		result = "down"
		break
	case "FF":
		result = "invalid"
		break
	}
	return result + "℃"
}

// led
func airConditionerLED(data string) string {
	result := "关"
	res := common.BinDec(data) //
	switch res {
	case 0:
		result = "开"
		break
	}
	return result
}

// cmd
func airConditionerCMD(data string) string {
	result := "开关"
	res := common.BinDec(data) //
	switch res {
	case 1:
		result = "非开关"
		break
	}
	return result
}

// 开关
func airConditionerType(data string) string {
	result := "reserve" // 00: 无状态; 01: 有状态; 02: 协议; 03: 推荐场景; 04: 半状态; 11: 忽略	空调类型
	res := common.BinHex(data)
	switch res {
	case "00":
		result = "无状态"
		break
	case "01":
		result = "有状态"
		break
	case "02":
		result = "协议"
		break
	case "03":
		result = "推荐场景"
		break
	case "04":
		result = "半状态"
		break
	case "11":
		result = "忽略"
		break
	}
	return result
}

// 转换空调参数
func AirConditionerToMap(data int) map[string]string {
	result := make(map[string]string, 0)
	binData := common.DecToBin(data) // 十进制转二进制
	result["switch"] = airConditionerSwitch(binData[0:4])
	result["model"] = airConditionerModel(binData[4:8])
	result["speed"] = airConditionerWindSpeed(binData[8:12])
	result["wind"] = airConditionerWindDirection(binData[12:14])
	result["sweeping"] = airConditionerSweeping(binData[14:16])
	result["temperature"] = airConditionerTemperature(binData[16:24])
	result["led"] = airConditionerLED(binData[26:27])
	result["cmd"] = airConditionerCMD(binData[27:28])
	return result
}

// 转换空调参数
func AirConditionerValueToMap(data int) map[string]interface{} {
	result := make(map[string]interface{}, 0)
	binData := common.DecToBin(data) // 十进制转二进制
	result["switch"] = airConditionerSwitch(binData[0:4])
	result["model"] = airConditionerModel(binData[4:8])
	result["speed"] = airConditionerWindSpeed(binData[8:12])
	result["wind"] = airConditionerWindDirection(binData[12:14])
	result["sweeping"] = airConditionerSweeping(binData[14:16])
	result["temperature"] = airConditionerTemperature(binData[16:24])
	result["led"] = airConditionerLED(binData[26:27])
	result["cmd"] = airConditionerCMD(binData[27:28])

	result["switch_val"] = common.BinHex(binData[0:4])
	result["model_val"] = common.BinHex(binData[4:8])
	result["speed_val"] = common.BinHex(binData[8:12])
	result["wind_val"] = common.BinHex(binData[12:14])
	result["sweeping_val"] = common.BinHex(binData[14:16])
	result["temperature_val"] = common.BinDec(binData[16:24])
	result["led_val"] = common.BinHex(binData[26:27])
	result["cmd_val"] = common.BinHex(binData[27:28])

	//result["类型"] = airConditionerType(binData[28:32])
	//fmt.Println(binData[24:25]) // 默认为0	扩展位
	//fmt.Println(binData[25:26]) // 默认为0	是否为压缩码
	return result
}

// 转换空调参数
func AirConditionerValueBinToMap(binData string) map[string]interface{} {
	result := make(map[string]interface{}, 0)
	result["switch"] = airConditionerSwitch(binData[0:4])
	result["model"] = airConditionerModel(binData[4:8])
	result["speed"] = airConditionerWindSpeed(binData[8:12])
	result["wind"] = airConditionerWindDirection(binData[12:14])
	result["sweeping"] = airConditionerSweeping(binData[14:16])
	result["temperature"] = airConditionerTemperature(binData[16:24])
	result["led"] = airConditionerLED(binData[26:27])
	result["cmd"] = airConditionerCMD(binData[27:28])

	result["switch_val"] = common.BinHex(binData[0:4])
	result["model_val"] = common.BinHex(binData[4:8])
	result["speed_val"] = common.BinHex(binData[8:12])
	result["wind_val"] = common.BinHex(binData[12:14])
	result["sweeping_val"] = common.BinHex(binData[14:16])
	result["temperature_val"] = common.BinDec(binData[16:24])
	result["led_val"] = common.BinHex(binData[26:27])
	result["cmd_val"] = common.BinHex(binData[27:28])

	//result["类型"] = airConditionerType(binData[28:32])
	//fmt.Println(binData[24:25]) // 默认为0	扩展位
	//fmt.Println(binData[25:26]) // 默认为0	是否为压缩码
	return result
}
