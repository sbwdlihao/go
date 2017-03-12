package basic

import (
	"testing"
	"fmt"
)

type contact struct {
	phone string
	address struct{
		city string
		street string
	     }
}

func TestNestStruct(t *testing.T) {
	c := contact{
		phone: "1234567890",
		address:struct {
			city string
			street string
		}{
			city: "hangzhou",
			street: "wensanlu",
		},
	}
	fmt.Printf("%v\n", c)
	// 输出 {1234567890 {hangzhou wensanlu}}

	//
	// 下面这种写法会报missing type in composite literal，也就是说address没有指定类型

	//c1 := contact{
	//	phone: "1234567890",
	//	address: {
	//		city: "hangzhou",
	//		street: "wensanlu",
	//	},
	//}
	//fmt.Printf("%v\n", c1)
}
