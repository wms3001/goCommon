package goCommon

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestSkuCode_Create(t *testing.T) {
	skuCode := SkuCode{}
	skuCode.Perfix = "YNHC"
	sku := skuCode.Create()
	fmt.Println(sku)

	b := make([]byte, 100)
	rand.Read(b)
	fmt.Println(b[0])
}
