package util

import "github.com/947133297/lwjmap/myMap"

var codeMap *myMap.MyMap

func init() {
	codeMap = myMap.CreateMapManager(3)

	codeMap.Set("1", "ydpm")
	codeMap.Set("2", "rnvo")
	codeMap.Set("3", "xu24")
	codeMap.Set("4", "muu4")
	codeMap.Set("5", "xeaw")
	codeMap.Set("6", "cwmx")
	codeMap.Set("7", "tedf")
	codeMap.Set("8", "2q5l")
	codeMap.Set("9", "uh32")
	codeMap.Set("10", "xapk")
}

func VerifyCode(key, code string) bool {
	v, ok := codeMap.Get(key)
	if !ok {
		return false
	}
	c, ok := v.(string)
	return ok && c == code
}
