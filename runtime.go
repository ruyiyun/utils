package utils

import (
	"encoding/json"
	"fmt"
	"runtime"
)

//返回调用此方法的方法名及key值,组合为缓存用的Key
func GetExecFuncId(key ...interface{}) (Id string) {

	if funcName, _, _, ok := runtime.Caller(1); ok {
		if len(key) > 0 {
			js1, _ := json.Marshal(key[0])

			Id = fmt.Sprintf("%v_%v", runtime.FuncForPC(funcName).Name(), MD5(string(js1)))
		} else {
			Id = fmt.Sprintf("%v", runtime.FuncForPC(funcName).Name())

		}
	}

	return
}
