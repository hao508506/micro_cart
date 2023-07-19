package common

import "encoding/json"

// 通过json tag进行结构体赋值
func SwapTo(req, cate interface{}) (err error) {
	dataByte, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, cate)
}
