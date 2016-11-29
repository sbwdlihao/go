package chain

import (
	"testing"
	"chain/crypto/ed25519/chainkd"
	"bytes"
	"io"
	"fmt"
)

func TestCrypto0(t *testing.T) {
	entrypy := [32]byte{0}
	xprv, _ := chainkd.NewXPrv(io.Reader(bytes.NewReader(entrypy[:])))
	fmt.Println(xprv[0], xprv[31]) // 120 104
	xpub := xprv.XPub()
	fmt.Println(xpub[0], xpub[63]) // 218 242
}

