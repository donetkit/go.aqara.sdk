package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 生成JSON字符串
func MarshalJson(m interface{}) string {
	str, _ := json.Marshal(m)
	return string(str)
}

func UnmarshalJson(data []byte, v interface{}) {
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

// json转map函数，通用
func JSONToMapString(str string) map[string]string {
	var tempMap map[string]string
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return tempMap
}

func StructToUrlParams(obj interface{}) string {
	var url = ""
	obj_v := reflect.ValueOf(obj)
	v := obj_v.Elem()
	typeOfType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		url += fmt.Sprintf("%s=%v&", typeOfType.Field(i).Tag.Get("json"), v.Field(i))
	}
	return url[:len(url)-1]
}
