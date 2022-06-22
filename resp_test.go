package goCommon

import (
	"fmt"
	"testing"
)

func TestResp_ToJson(t *testing.T) {
	resp := Resp{}
	resp.Code = 1
	resp.Message = "This is test"
	resp.Data = `{"Code":1,"Message":"connected","Data":""}`
	fmt.Println(resp.ToJson())
}
