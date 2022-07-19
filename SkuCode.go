package goCommon

import (
	"crypto/rand"
	"fmt"
	"time"
)

type SkuCode struct {
	Perfix string `json:"perfix"`
}

func (skuCode *SkuCode) Create() string {
	b := make([]byte, 100)
	rand.Read(b)
	return skuCode.Perfix + fmt.Sprint(time.Now().Unix()) + fmt.Sprint(b[0])
}
