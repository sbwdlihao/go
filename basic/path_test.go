package basic

import (
	"testing"
	"path/filepath"
	"fmt"
)

func TestAbs(t *testing.T) {
	abs, _ := filepath.Abs("fmt_test.go")
	// 运行单元测试的情况下输出 /Users/sbwdlihao/GoProjects/src/github.com/sbwdlihao/go/basic/fmt_test.go
	// 不过如果是在运行main的情况下不一样，比如web/sample/sample.go中的例子
	fmt.Println(abs)
}