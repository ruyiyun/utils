package utils

import (
	"fmt"
	"runtime"
)

//返回调用此方法的方法名及key值,组合为缓存用的Key
func GetExecFuncId(key interface{}) (Id string) {
	if funcName, _, _, ok := runtime.Caller(1); ok {
		Id = fmt.Sprintf("%v_%v", runtime.FuncForPC(funcName).Name(), key)
	}
	return
}
