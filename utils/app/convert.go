package app

import (
	"encoding/json"
	"fmt"
)

// json转struct
func JsonToStruct(jsonStr string, obj interface{}) {
	Error("json转struct", json.Unmarshal([]byte(jsonStr), &obj))
}

// struct转json
func StructToJson(obj interface{}) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonBytes)
}

// json转map
func JsonToMap(jsonStr string) (result map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return result
	}
	return result
}

// map转json
func MapToJson(instance map[string]interface{}) string {
	jsonStr, err := json.Marshal(instance)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	return string(jsonStr)
}
