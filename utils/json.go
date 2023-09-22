/**
 * @Auth: Nuts
 * @Date: 2021/3/31 10:30 上午
 */
package utils

import (
	"encoding/json"
	"reflect"
)

type JsonUtils struct {
}

// json 转码，转成 string
func (JsonUtils) JsonEncode(v interface{}) string {
	res, _ := JsonUtils{}.JsonEncodeBytes(v)
	return string(res)
}

func (JsonUtils) JsonEncodeBytes(v interface{}) ([]byte, error) {
	res, err := json.Marshal(v)
	return res, err
}

// json 解码，
func (JsonUtils) JsonDecode(s string, v interface{}) error {
	return JsonUtils{}.JsonDecodeBytes([]byte(s), v)
}

// json 解码，需要注意数字会被解码成 float64
func (JsonUtils) JsonDecodeBytes(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

// json解码成map
func (JsonUtils) JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Convert map json string
func MapToJson(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		return "", nil
	}

	return string(jsonByte), nil
}

func StructToJsonTagMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	data := make(map[string]interface{})

	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		key := objType.Field(i).Tag.Get("json")
		value := field.Interface()
		data[key] = value
	}

	return data
}
