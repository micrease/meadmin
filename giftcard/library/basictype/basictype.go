package basictype

import (
	"github.com/shopspring/decimal"
	"math/rand"
	"time"
)

//decimal
func DecimalFloor(n decimal.Decimal, places int32) string {
	val := n.Truncate(places).StringFixedBank(places)
	return val
}

//生成随机字符串
func GetRandomNum(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
