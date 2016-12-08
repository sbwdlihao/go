package btcsuite

import (
	"testing"
	"github.com/btcsuite/btcd/txscript"
	"fmt"
	"encoding/hex"
)

func TestScript0(t *testing.T) {
	sb := txscript.NewScriptBuilder()
	sb.AddInt64(442465)
	sb.AddInt64(1480406530)
	info := []byte("/BTC.COM/")
	fmt.Println(info)
	sb.AddFullData(info)
	b, _ := sb.Script()
	fmt.Println(hex.EncodeToString(b)) // 0361c006 0402363d58
}
