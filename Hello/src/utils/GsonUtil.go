package utils

import (
	"encoding/json"
)

//空接口可以作为任何类型使用
type AnyThing interface {
}

//结构体 转 JSON string
func EntityToJson(at AnyThing) string {
	data, _ := json.Marshal(at)
	return string(data)
}
