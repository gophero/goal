package mempool_test

import (
	"github.com/gophero/goal/btcx"
	"github.com/gophero/goal/btcx/mempool"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func TestGetRawTransaction(t *testing.T) {
	//https://mempool.space/signet/api/tx/b752d80e97196582fd02303f76b4b886c222070323fb7ccd425f6c89f5445f6c/hex
	client := mempool.NewClient(btcx.TestNet)
	txId, _ := chainhash.NewHashFromStr("b752d80e97196582fd02303f76b4b886c222070323fb7ccd425f6c89f5445f6c")
	transaction, err := client.GetRawTransaction(txId)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(transaction.TxHash().String())
	}
}
