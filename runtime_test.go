package utils

import (
	"testing"
)

func TestName(t *testing.T) {
	key := "test"
	funcId := GetExecFuncId(key)
	if funcId != "github.com/ruyiyun/utils.TestName_test" {
		t.Fail()
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetExecFuncId(i)
	}
}
