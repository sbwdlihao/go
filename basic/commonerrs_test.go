package basic

import (
	"testing"
	"fmt"
	"time"
)

// 测试for range变量的时候，变量会在不同的go routines之间重用
func TestForRangeVar(t *testing.T)  {
	bs := []byte{0, 1, 2, 3}

	for _, b := range bs {
		go func() {
			fmt.Println(b)
		}()
	} // 全部输出的都是3

	time.Sleep(time.Second)

	for _, b := range bs {
		b := b
		go func() {
			fmt.Println(b)
		}()
	} // 输出的0,1,2,3

	time.Sleep(time.Second)

	for _, b := range bs {
		go func(b byte) {
			fmt.Println(b)
		}(b)
	} // 输出的0,1,2,3
}
