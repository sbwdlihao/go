package basic

import (
	"testing"
	"encoding/json"
	"os"
	"fmt"
)

type template struct {
	CoinbaseValue int64 `json:"coinbasevalue,omitempty"`
}

func TestJson0(t *testing.T) {
	t0 := template{
		CoinbaseValue: 5000017380,
	}
	bytes, err := json.Marshal(t0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var t1 template
	err = json.Unmarshal(bytes, &t1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(t1.CoinbaseValue)
}

func TestJson1(t *testing.T)  {
	m := map[string]string{
		"a": "da",
		"b": "d",
	}
	mb, _ := json.Marshal(m)

	var mc map[string]string
	//mc := make(map[string]string) 这样写也可以
	json.Unmarshal(mb, &mc)
	fmt.Println(mc)
}

func TestString(t *testing.T) {
	var s string
	s = "7f6"
	b, _ := json.Marshal(s)
	fmt.Println(b) // [34 55 102 54 34]
	var s1 string
	json.Unmarshal(b, &s1)
	fmt.Println(s1) // 7f6
	parseString(&s1) // abd
	fmt.Println(s1)
}

func parseString(v interface{})  {
	json.Unmarshal([]byte{34, 97, 98, 100, 34}, v)
}

