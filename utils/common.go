package utils

import (
	"encoding/json"
	"go.uber.org/zap"
	"im_app/global"
)

func Marshal(m map[string]interface{}) string {
	if byt, err := json.Marshal(m); err != nil {
		return ""
	} else {
		return string(byt)
	}
}

func Unmarshal(str string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

// map to
func MapToStruct(data any, msg any) (res any, err error) {
	// 反序列化
	arr, err := json.Marshal(data)
	if err != nil {
		global.GVA_LOG.Error("json.Marshal(data)", zap.Error(err))
		return
	}
	// 反序列化
	err = json.Unmarshal(arr, &msg)
	if err != nil {
		global.GVA_LOG.Error("json.Unmarshal(arr, &msg)", zap.Error(err))
		return
	}

	return msg, nil

}
