package util

import (
	"math/rand"
	"strconv"

	"github.com/947133297/lwjmap/myMap"
)

var codeMap *myMap.MyMap

func init() {
	codeMap = myMap.CreateMapManager(5)
}

func GetCode() (key, code string) {
	code = strconv.Itoa(rand.Intn(8999) + 1000)
	key = strconv.Itoa(rand.Intn(9999))
	codeMap.Set(key, code)
	return
}

func VerifyCode(key, code string) bool {
	v, ok := codeMap.Get(key)
	if !ok {
		return false
	}
	c, ok := v.(string)
	pass := ok && c == code
	if pass {
		codeMap.Delete(key)
	}
	return pass
}
