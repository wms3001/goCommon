package goCommon

import "encoding/json"

//
//  Resp
//  @Description: 统一返回结果
//
type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

//
//  ToJson
//  @Description: 返回json
//  @receiver resp
//  @return string
//
func (resp *Resp) ToJson() string {
	json, _ := json.Marshal(resp)
	return string(json)
}
