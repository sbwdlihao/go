package btcsuite

import (
	"testing"
	"github.com/btcsuite/btcd/wire"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"encoding/hex"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"bytes"
	"github.com/btcsuite/btcd/txscript"
)

func TestTx0(t *testing.T) {
	var tx *wire.MsgTx = &wire.MsgTx{}
	var tx1 *wire.MsgTx = new(wire.MsgTx)
	fmt.Println(*tx, *tx1) // {0 [] [] 0} {0 [] [] 0}

	// coinbase txin
	txInPreOutPoint := wire.NewOutPoint(&chainhash.Hash{}, wire.MaxTxInSequenceNum)
	txInScript, _ := hex.DecodeString("0359c0060492dd48582f4254432e434f4d2ffabe6d6d1d1afd7870501b345070d79f98e0c30b151efb9d9ca533a8aa1e4bbf32dd8c590100000000000000250001d27a11005800000000")
	var txIn *wire.TxIn = wire.NewTxIn(txInPreOutPoint, txInScript)
	amt, _ := btcutil.NewAmount(12.62259455)
	address, _ := btcutil.DecodeAddress("3NA8hsjfdgVkmmVS9moHmkZsVCoLxUkvvv", &chaincfg.MainNetParams)
	pkScript, _ := txscript.PayToAddrScript(address)
	var txOut *wire.TxOut = wire.NewTxOut(int64(amt), pkScript)
	tx.Version = wire.TxVersion
	tx.LockTime = 0
	tx.AddTxIn(txIn)
	tx.AddTxOut(txOut)

	var bf bytes.Buffer
	tx.Serialize(&bf)
	fmt.Println(tx.TxHash()) // 0784c3f0ec72854e723bad27c613820dd3c423d3615cc4e54204731c00148f91
	fmt.Println(hex.EncodeToString(bf.Bytes()))
	// 01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4a0359c0060492dd48582f4254432e434f4d2ffabe6d6d1d1afd7870501b345070d79f98e0c30b151efb9d9ca533a8aa1e4bbf32dd8c590100000000000000250001d27a11005800000000ffffffff01ff8c3c4b0000000017a914e083685a1097ce1ea9e91987ab9e94eae33d8a138700000000
}
