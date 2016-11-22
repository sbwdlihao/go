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

func Test0(t *testing.T) {
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

