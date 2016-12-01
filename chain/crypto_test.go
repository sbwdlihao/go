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

	sel := [2]byte{0,1}
	child := xprv.Child(sel[:], true)
	fmt.Println(child[0], child[63]) // 16 201
	child = xprv.Child(sel[:], false)
	fmt.Println(child[0], child[63]) // 36 10

	path := [][]byte{
		{0,1},
		{3,4},
	}
	derive := xprv.Derive(path)
	fmt.Println(derive[0], derive[63]) // 246 162

	pub_child := xpub.Child(sel[:]);
	fmt.Println(pub_child[0], pub_child[63]); // 60 10

	pub_derive := xpub.Derive(path)
	fmt.Println(pub_derive[0], pub_derive[63]) // 103 162

	msg := []byte{0,1}
	sign := xprv.Sign(msg)
	fmt.Println(sign[0], sign[63]) // 173, 11

	xprv_text, _ := xprv.MarshalText()
	xpub_text, _ := xpub.MarshalText()
	fmt.Println(string(xprv_text)) // 78ce8a0989f5cebfefbee1bdcae71d4ae2757d9c7e0e8a404444774456d1bc681c199d4171a46bf486704c23522c2f962468afaffee0d6afc1fb84c1e4433ff2
	fmt.Println(string(xpub_text)) // da5ddf057de3cf8db186e619915b88f37f797a9be1a5a79195533741f8bec4a21c199d4171a46bf486704c23522c2f962468afaffee0d6afc1fb84c1e4433ff2
	fmt.Println(xprv.String()) // 78ce8a0989f5cebfefbee1bdcae71d4ae2757d9c7e0e8a404444774456d1bc681c199d4171a46bf486704c23522c2f962468afaffee0d6afc1fb84c1e4433ff2
}

