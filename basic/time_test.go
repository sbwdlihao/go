package basic

import (
	"testing"
	"fmt"
	"time"
	"encoding/json"
)

func TestTime0(t *testing.T) {
	a := 1
	//b := a * time.Second // invalid operation: a * time.Second (mismatched types int and time.Duration)
	b := time.Duration(a) * time.Second // ok
	fmt.Println(b)
}

func TestTime1(t *testing.T) {
	s, _ := json.Marshal(map[string]string{
		"time": time.Now().Format(time.RFC3339),
	})
	fmt.Println(string(s))

	var j map[string]interface{}
	json.Unmarshal(s, &j)
	ts := j["time"].(string)
	fmt.Println(time.Parse(time.RFC3339, ts))
}
