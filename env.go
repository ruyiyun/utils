package utils

import "os"

//获取系统变量,如果没有设置就用默认值
func GetEnv(key string, def ...string) string {

	if v := os.Getenv(key); v != "" {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}
	return ""

}
