package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"testing"

	"github.com/donetkit/go.aqara.sdk/app/device_utils"

	"github.com/donetkit/go.aqara.sdk/app/frame/util"

	"github.com/donetkit/go.aqara.sdk/app/device_utils/common"

	"github.com/donetkit/go.aqara.sdk/app/config"
	"github.com/donetkit/go.aqara.sdk/app/device_sign"
	"github.com/donetkit/go.aqara.sdk/app/device_token"
	"github.com/donetkit/go.aqara.sdk/app/frame/auth_client"
	"github.com/donetkit/go.aqara.sdk/app/frame/dev_client"
	"github.com/donetkit/go.aqara.sdk/app/frame/gb"
)

var did = "lumi."

func InitData() {
	config.SetupRedisConfig("../config/redis.yml")
	config.SetupAqaraConfig("../config/aqara.yml")
	gb.RegisterCache()
	auth_client.RegisterAuthClient()
	dev_client.RegisterDevClient()
}

// 查询设备
func TestTokenDeviceQueryDid(t *testing.T) {
	InitData()
	t.Log("testing -> DeviceQueryDid")
	// 查询设备
	result := device_token.ApiDevice.DeviceQueryDid(did, 1, 50)
	t.Log(result)
}

// 查询设备
func TestTokenDeviceQuery(t *testing.T) {
	InitData()
	t.Log("testing -> DeviceQuery")
	// 查询设备
	result := device_token.ApiDevice.DeviceQuery()
	t.Log(result)
}

// 查询设备
func TestTokenDevicePageQuery(t *testing.T) {
	InitData()
	t.Log("testing -> DeviceQueryPage")
	// 查询设备
	var pageNum = 0
	for {
		pageNum++
		data := device_token.ApiDevice.DeviceQueryPage(pageNum, 50)
		t.Log(data)
		if len(data) < 50 {
			return
		}

	}

}

// 查询子设备
func TestTokenDeviceChildQuery(t *testing.T) {
	InitData()
	t.Log("testing -> DeviceChildQuery")
	device_token.ApiDevice.DeviceChildQuery(did)

}

// 修改设备名称
func TestTokenUpdateDeviceName(t *testing.T) {
	InitData()
	t.Log("testing -> UpdateDeviceName")
	result := device_token.ApiDevice.UpdateDeviceName(did, "二年级(1)班")
	t.Log(result)
}

// 查询设备属性
func TestSignDevicePropertiesQuery(t *testing.T) {
	InitData()
	t.Log("testing -> TestSignDevicePropertiesQuery")
	result := device_sign.ApiSignProperties.PropertiesQuery(did, "relay_status")

	//result := device_sign.ApiSignProperties.PropertiesQuery([]string{"lumi.", "lumi."}, "ac_state")

	t.Log(result)
}

// 查询设备（温度）属性
func TestSignDevicePropertiesAllQuery(t *testing.T) {
	InitData()
	t.Log("testing -> PropertiesAllQuery")
	//result := device_sign.ApiSignProperties.PropertiesAllQuery("lumi.")
}

// 控制设备
func TestSignDevicePropertiesWrite(t *testing.T) {
	InitData()
	t.Log("testing -> PropertiesWrite")

	//result := device_sign.ApiSignProperties.PropertiesWrite(did, "", "")
	//t.Log(result)
}

// 查询设备历史记录
func TestSignDevicePropertiesHistoryDateTimeQuery(t *testing.T) {
	InitData()
	t.Log("testing -> TestSignDevicePropertiesHistoryDateTimeQuery")
	result := device_sign.ApiSignProperties.PropertiesHistoryDateTimeQuery(did, 1612287085000, 1612315885000, 10)
	t.Log(result)
}

// 查询设备历史记录
func TestSignDevicePropertiesHistoryQuery(t *testing.T) {
	InitData()
	t.Log("testing -> TestSignDevicePropertiesHistoryQuery")
	result := device_sign.ApiSignProperties.PropertiesHistoryQuery(did, []string{"relay_status", "ac_state", "ac_load_power"}, 10)
	t.Log(result)
}

// 打开空调
func TestOpenDevice(t *testing.T) {
	InitData()
	t.Log("testing -> PropertiesWrite")

	result := device_sign.ApiSignProperties.PropertiesQuery(did, "ac_state")
	t.Log(result)

	if len(result) > 0 && result[0].Property == "ac_state" {
		data, _ := strconv.Atoi(result[0].Value)
		binData := common.DecToBin(data)
		//t.Log(binData)
		//t.Log(binData[4:])
		cmd := "0001" + binData[4:]
		//t.Log(cmd)
		//00000000000000000001100000000010
		//0000000000000001100000000010

		var cmdopen = fmt.Sprintf("%d", common.BinDec(cmd))
		t.Log("==============================================")
		t.Log(cmdopen)
		result1 := device_sign.ApiSignProperties.PropertiesWrite(did, "ac_state", "00010010000000000001010100000001")
		t.Log(result1)
	}

}

// 关闭空调
func TestCloseDevice(t *testing.T) {
	InitData()
	t.Log("testing -> PropertiesWrite")

	result := device_sign.ApiSignProperties.PropertiesQuery(did, "ac_state")
	t.Log(result)

	if len(result) > 0 && result[0].Property == "ac_state" {
		data, _ := strconv.Atoi(result[0].Value)
		binData := common.DecToBin(data)
		//t.Log(binData)
		//t.Log(binData[4:])
		cmd := "0000" + binData[4:]
		//t.Log(cmd)
		//00000000000000000001100000000010
		//0000000000000001100000000010

		var cmdopen = fmt.Sprintf("%d", common.BinDec(cmd))
		t.Log("==============================================")
		t.Log(cmdopen)
		result1 := device_sign.ApiSignProperties.PropertiesWrite(did, "ac_state", cmdopen)
		t.Log(result1)
	}
}

// 授权测试
func TestAuthorizeDevice(t *testing.T) {
	InitData()
	t.Log("testing -> OauthAuthorize2")
	//device_auth.ApiAuthAuthorize.OauthAuthorize2(nil)
	//result := device_auth.ApiAuthAuthorize.OauthAuthorize2
	//t.Log(result)
	//
	//if len(result) > 0 && result[0].Property == "ac_state" {
	//	data, _ := strconv.Atoi(result[0].Value)
	//	binData := common.DecToBin(data)
	//	//t.Log(binData)
	//	//t.Log(binData[4:])
	//	cmd := "0000" + binData[4:]
	//	//t.Log(cmd)
	//	//00000000000000000001100000000010
	//	//0000000000000001100000000010
	//
	//	var cmdopen = fmt.Sprintf("%d", common.BinDec(cmd))
	//	t.Log("==============================================")
	//	t.Log(cmdopen)
	//	result1 := device_sign.ApiSignProperties.PropertiesWrite(did, "ac_state", cmdopen)
	//	t.Log(result1)
	//}
}

// 关闭空调
func TestP3Device(t *testing.T) {
	InitData()
	t.Log("testing -> PropertiesWrite") // P0_M2_T21_S0_L1
	// 协议码空调、有状态空调格式Px_Mm_Ty_Ss_Dd。如：空调需要改变模式到制冷，温度到18度，这参数为"M0_T18"。 * x 表示开关 (AC_POWER_ON = 0, AC_POWER_OFF = 1) * m 表示模式(0:制冷,1:制热,2:自动,3:送风,4:除湿) * y 表示温度(温度范围一般为16~30) * s 表示风速( 0:自动,1:小风量,2:中风量,3:大风量) * d 表示风向(0 表示扫风,其他值表示固定风向) * L 表示灯光(0关灯光 1开灯光) 无状态空调，红外热水器参数为按键id，如"1"。

	resultac := device_sign.ApiSignProperties.PropertiesQuery("lumi1.", "send_ac_cmd")
	t.Log(resultac)
	result := device_sign.ApiSignProperties.PropertiesWrite("lumi1.", "send_ac_cmd", "P0_M2_T21_S0_L1")
	t.Log(result)
	//
	//if len(result) > 0 && result[0].Property == "ac_state" {
	//	data, _ := strconv.Atoi(result[0].Value)
	//	binData := common.DecToBin(data)
	//	//t.Log(binData)
	//	//t.Log(binData[4:])
	//	cmd := "0000" + binData[4:]
	//	//t.Log(cmd)
	//	//00000000000000000001100000000010
	//	//0000000000000001100000000010
	//
	//	var cmdopen = fmt.Sprintf("%d", common.BinDec(cmd))
	//	t.Log("==============================================")
	//	t.Log(cmdopen)
	//	result1 := device_sign.ApiSignProperties.PropertiesWrite(did, "ac_state", cmdopen)
	//	t.Log(result1)
	//}
}

// 设备 系统音量，取值范围为0~100。
func TestDeviceSystemVolume(t *testing.T) {
	InitData()
	t.Log("testing -> PropertiesWrite")

	result := device_sign.ApiSignProperties.PropertiesQuerySystemVolume(did)
	t.Log(result)

	result1 := device_sign.ApiSignProperties.PropertiesWriteSystemVolume(did, 40)
	t.Log(result1)

}

// 设备 继电器开关状态。0:关闭，1:打开,2:unknown(读)
func TestDeviceRelayStatus(t *testing.T) {
	InitData()
	t.Log("testing -> TestDeviceRelayStatus")

	result := device_sign.ApiSignProperties.PropertiesQueryRelayStatus(did)
	t.Log(result)

	//result1 := device_sign.ApiSignProperties.PropertiesWriteRelayStatus(did, 1)
	//t.Log(result1)

}

// 设备 空调处于打开/关闭状态,0:关闭，1:打开。
func TestDeviceOnOffStatus(t *testing.T) {
	InitData()
	t.Log("testing -> TestDeviceOnOffStatus")

	result := device_sign.ApiSignProperties.PropertiesQueryOnOffStatus(did)
	t.Log(result)

}

// 设备 空调处于打开/关闭状态,0:关闭，1:打开。
func TestDeviceState(t *testing.T) {
	InitData()
	//result1 := device_sign.ApiSignProperties.PropertiesWrite(did, "ac_state", "00000000000000000001011000000001")
	//t.Log(result1)
	//
	//result1 = device_sign.ApiSignProperties.PropertiesWrite(did, "ac_state", "00010000000000000001011000000001")
	//t.Log(result1)

	t.Log(device_utils.AirConditionerToMap(268441089))

	binData := common.DecToBin(268441089)
	t.Log(binData)
	//fmt.Println(common.BinHex(binData[0:4]))
	//t.Log(common.BinHex(binData[0:4]))
	//t.Log(device_utils.AirConditionerValueToMap(301995265))

	// 00010001000000000001010100000001
	// 00010001000000000001010100000001
	// 00000001000000000001010100000001

	//t.Log(common.BinDec("00000001000000000001010100000001"))
	////t.Log(common.BinDec("00010000000000010001011100000001"))
	//t.Log(device_utils.AirConditionerValueToMap(16782593))
	//t.Log(common.BinDec("00000010000000000001011100000001"))
	//t.Log(common.BinDec("00010001000000000001010100000001"))

}

// 设备 空调处于打开/关闭状态,0:关闭，1:打开。
func TestDeviceCH(t *testing.T) {
	InitData()
	t.Log(device_sign.ApiSignProperties.PropertiesQueryOnOffStatus("lumi."))
	t.Log(device_sign.ApiSignProperties.PropertiesAllQuery("lumi."))
}

// AES
func TestAES(t *testing.T) {
	payload, _ := util.Encrypt([]byte("123456"), util.GetAesIv([]byte("PsVLSHNRrBdlE7BAAMdnmKJP8rr3g9Rp")))

	fmt.Println(base64.StdEncoding.EncodeToString(payload))
	temp, _ := util.Decryptbase64DecodeString(base64.StdEncoding.EncodeToString(payload), []byte("PsVLSHNRrBdlE7BAAMdnmKJP8rr3g9Rp"))
	fmt.Println(string(temp))
	//fmt.Print(base64.StdEncoding.EncodeToString(temp))
}
