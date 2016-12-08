package btcsuite

import (
	"testing"
	"github.com/btcsuite/btcd/blockchain"
	"fmt"
	"strconv"
)

func TestBasic0(t *testing.T) {
	target := blockchain.CompactToBig(402904457)
	bits := blockchain.BigToCompact(target)
	fmt.Println(bits) // 402904457
	bitss := strconv.FormatInt(int64(bits), 16)
	fmt.Println(bitss) // 1803d589
	bitsi, _ := strconv.ParseInt(bitss, 16, 64)
	fmt.Println(bitsi) // 402904457
	target1 := blockchain.CompactToBig(uint32(bitsi))
	fmt.Println(target1) // 94012390634764280055243391736606357298689315295029362688
}
