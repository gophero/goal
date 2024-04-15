package mempool

import (
	"encoding/json"
	"fmt"
	"github.com/gophero/goal/btcx"
	"github.com/gophero/goal/btcx/mempool/model"
	"github.com/gophero/goal/btcx/util"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

func (c *MempoolClient) GetAddress(addr string) (model.AddressInfo, error) {
	if r, err := c.get(fmt.Sprintf("/address/%s", addr)); err != nil {
		return model.AddressInfo{}, err
	} else {
		var v model.AddressInfo
		return v, json.Unmarshal(r, &v)
	}
}

// ListTxs Get transaction history for the specified address/scripthash, sorted
// with newest first. Returns up to 50 mempool transactions plus the first 25
// confirmed transactions. You can request more confirmed transactions using an
// after_txid query parameter.
func (c *MempoolClient) ListTxs(addr string, afterTxid string) ([]*wire.MsgTx, error) {
	res, err := c.get(fmt.Sprintf("/address/%s/txs?after_txid=%s", addr, afterTxid))
	if err != nil {
		return nil, err
	}

	var msgTxs []*wire.MsgTx
	var txs []*model.Tx
	err = json.Unmarshal(res, &txs)
	for _, v := range txs {
		msgTxs = append(msgTxs, v.MsgTx(c.Net))
	}
	return msgTxs, err
}

// ListTxsConfirmed get confirmed transaction history for the specified
// address/scripthash, sorted with newest first. Returns 25 transactions per
// page. More can be requested by specifying the last txid seen by the previous
// query.
func (c *MempoolClient) ListTxsConfirmed(addr string, lastTxid string) ([]*wire.MsgTx, error) {
	res, err := c.get(fmt.Sprintf("/address/%s/txs/chain/%s", addr, lastTxid))
	if err != nil {
		return nil, err
	}

	var msgTxs []*wire.MsgTx
	var txs []*model.Tx
	err = json.Unmarshal(res, &txs)
	for _, v := range txs {
		msgTxs = append(msgTxs, v.MsgTx(c.Net))
	}
	return msgTxs, err
}

func (c *MempoolClient) ListUnspent(addr string) ([]*btcx.UnspentOutput, error) {
	address, err := util.ParseAddress(addr, c.Net)
	if err != nil {
		return nil, err
	}
	res, err := c.get(fmt.Sprintf("/address/%s/utxo", addr))
	if err != nil {
		return nil, err
	}

	var utxos []*model.UTXO
	err = json.Unmarshal(res, &utxos)
	if err != nil {
		return nil, err
	}

	unspentOutputs := make([]*btcx.UnspentOutput, 0)
	for _, utxo := range utxos {
		txHash, err := chainhash.NewHashFromStr(utxo.Txid)
		if err != nil {
			return nil, err
		}
		unspentOutputs = append(unspentOutputs, &btcx.UnspentOutput{
			Outpoint: wire.NewOutPoint(txHash, uint32(utxo.Vout)),
			Output:   wire.NewTxOut(utxo.Value, address.ScriptAddress()),
		})
	}
	return unspentOutputs, nil
}

// ListTxsUnconfirmed Get unconfirmed transaction history for the specified address/scripthash.
// Returns up to 50 transactions (no paging).
func (c *MempoolClient) ListTxsUnconfirmed(addr string) ([]*wire.MsgTx, error) {
	res, err := c.get(fmt.Sprintf("/address/%s/txs/mempool", addr))
	if err != nil {
		return nil, err
	}

	var msgTxs []*wire.MsgTx
	var txs []*model.Tx
	err = json.Unmarshal(res, &txs)
	for _, v := range txs {
		msgTxs = append(msgTxs, v.MsgTx(c.Net))
	}
	return msgTxs, err
}

func (c *MempoolClient) ValidAddress(addr string) (model.ValidAddress, error) {
	res, err := c.get(fmt.Sprintf("/v1/validate-address/%s", addr))
	if err != nil {
		return model.ValidAddress{}, err
	}
	var va model.ValidAddress
	return va, json.Unmarshal(res, &va)
}
