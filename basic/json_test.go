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

