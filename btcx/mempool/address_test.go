package mempool_test

import (
	"bytes"
	"encoding/hex"
	"github.com/gophero/goal/btcx"
	"github.com/gophero/goal/btcx/mempool"
	"github.com/gophero/goal/btcx/util"
	"testing"

	"github.com/btcsuite/btcd/txscript"
	"github.com/stretchr/testify/assert"
)

func TestGetAddress(t *testing.T) {
	client := mempool.NewClient(btcx.TestNet)
	r, err := client.GetAddress("tb1pja2cxxa3wmwpmce4jpdmhv3ulxhe9h2nqdglvv69z2c56eljhw9s3mn9tz")
	if err != nil {
		t.Error(err)
	}
	t.Logf("result: %v", r)
	assert.True(t, r.Address != "")
}

func TestListTxs(t *testing.T) {
	client := mempool.NewClient(btcx.TestNet)
	addr := "tb1pja2cxxa3wmwpmce4jpdmhv3ulxhe9h2nqdglvv69z2c56eljhw9s3mn9tz"
	txs, err := client.ListTxs(addr, "")
	if err != nil {
		panic(err)
	}
	assert.True(t, len(txs) > 0)
}

func TestListTxsConfirmed(t *testing.T) {
	client := mempool.NewClient(btcx.TestNet)
	addr := "tb1pja2cxxa3wmwpmce4jpdmhv3ulxhe9h2nqdglvv69z2c56eljhw9s3mn9tz"
	txs, err := client.ListTxsConfirmed(addr, "")
	if err != nil {
		t.Errorf("test failed: %v", err)
	}
	assert.True(t, len(txs) == 25)
	txs, err = client.ListTxsConfirmed(addr, "38a509320b0f6c21912cd6d55cfca29e3e53c452c28dd40e73ffe087fd646d02")
	if err != nil {
		t.Errorf("test failed: %v", err)
	}
	assert.True(t, len(txs) == 25)
	assert.True(t, txs[0].TxHash().String() == "e4b3c47128b1891f05ec2d27c7e9f67769f16b7f658c848ef642620f61df92ce")
	var firstTx = txs[0]
	var w bytes.Buffer
	if err := firstTx.Serialize(&w); err != nil {
		t.Errorf("serialize error: %v", err)
	}
	assert.True(t, firstTx.TxIn[0].PreviousOutPoint.Hash.String() == "245b347445fca06f7a8e8fc7d7a7d4c57936442b308c654ef679a1778bafae25")
	assert.True(t, firstTx.TxIn[0].PreviousOutPoint.Index == 1)
}

func TestListUnspent(t *testing.T) {
	// https://mempool.space/signet/api/address/tb1p8lh4np5824u48ppawq3numsm7rss0de4kkxry0z70dcfwwwn2fcspyyhc7/utxo
	client := mempool.NewClient(btcx.TestNet)
	addr := "tb1pja2cxxa3wmwpmce4jpdmhv3ulxhe9h2nqdglvv69z2c56eljhw9s3mn9tz"
	unspentList, err := client.ListUnspent(addr)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(len(unspentList))
		for _, output := range unspentList {
			t.Log(output.Outpoint.Hash.String(), "    ", output.Outpoint.Index)
			assert.True(t, output.Outpoint.Hash.String() != "")
		}
	}
}

func TestListUnconfirmed(t *testing.T) {
	client := mempool.NewClient(btcx.TestNet)
	addr := "tb1pja2cxxa3wmwpmce4jpdmhv3ulxhe9h2nqdglvv69z2c56eljhw9s3mn9tz"
	list, err := client.ListTxsUnconfirmed(addr)
	if err != nil {
		t.Error(err)
	} else {
		assert.True(t, len(list) == 0)
	}
}

func TestValidAddress(t *testing.T) {
	client := mempool.NewClient((btcx.TestNet))
	addr := "tb1pja2cxxa3wmwpmce4jpdmhv3ulxhe9h2nqdglvv69z2c56eljhw9s3mn9tz"
	va, err := client.ValidAddress(addr)
	if err != nil {
		t.Errorf("test failed: %v", err)
	} else {
		assert.True(t, va.Isvalid)
		assert.True(t, va.Iswitness)
		assert.True(t, va.Isscript)
		ad, _ := util.ParseAddress(addr, btcx.TestNet)
		pkScript, _ := txscript.PayToAddrScript(ad)
		pubscriptstr := hex.EncodeToString(pkScript)
		t.Log(pubscriptstr)
		assert.True(t, pubscriptstr == va.ScriptPubKey)
	}
}
