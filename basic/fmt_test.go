package basic

import (
	"testing"
	"fmt"
)

func wrapPrintln(v ...interface{})  {
	fmt.Println(v...)
}

func TestFmt0(t *testing.T) {
	wrapPrintln("a", "b")
}

