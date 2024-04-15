package model

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/wire"
	"github.com/gophero/goal/btcx"
	"github.com/gophero/goal/btcx/util"
	"github.com/gophero/goal/logx"
)

type UTXO struct {
	Txid   string      `json:"txid"`
	Vout   int         `json:"vout"`
	Status BlockStatus `json:"status"`
	Value  int64       `json:"value"`
}

type Tx struct {
	Txid     string       `json:"txid"`
	Version  int32        `json:"version"`
	LockTime uint32       `json:"locktime"`
	Vin      []*TxVin     `json:"vin"`
	Vout     []*TxVout    `json:"vout"`
	Size     int64        `json:"size"`
	Weight   int32        `json:"weight"`
	Sigops   int16        `json:"sigops"`
	Fee      int32        `json:"fee"`
	Status   *BlockStatus `json:"status"`
}

func (t *Tx) MsgTx(net btcx.Net) *wire.MsgTx {
	var msgTx = wire.NewMsgTx(t.Version)
	for _, v := range t.Vin {
		msgTx.AddTxIn(v.TxIn(net))
	}
	for _, v := range t.Vout {
		msgTx.AddTxOut(v.TxOut())
	}
	msgTx.LockTime = t.LockTime
	return msgTx
}

type TxVin struct {
	Txid         string   `json:"txid"`
	Vout         uint32   `json:"vout"`
	Scriptsig    string   `json:"scriptsig"`
	ScriptsigAsm string   `json:"scriptsig_asm"`
	IsCoinbase   bool     `json:"is_coinbase"`
	Sequence     uint32   `json:"sequence"`
	Prevout      TxVout   `json:"prevout"`
	Witness      []string `json:"witness"`
}

func (in *TxVin) TxIn(net btcx.Net) *wire.TxIn {
	var wti = &wire.TxIn{}
	hash, err := util.ParseTxHash(in.Txid)
	if err != nil {
		logx.Default.Errorf("decode tx %s error: %v\n", in.Txid, err)
		return nil
	}
	// TODO not sure
	sigScript, err := hex.DecodeString(in.Scriptsig)
	if err != nil {
		logx.Default.Errorf("decode sigScript error: %v\n", err)
		return nil
	}
	var witness [][]byte
	for _, w := range in.Witness {
		bs, err := hex.DecodeString(w)
		if err != nil {
			return nil
		}
		witness = append(witness, bs)
	}
	wti.PreviousOutPoint = wire.OutPoint{Hash: *hash, Index: in.Vout}
	wti.Sequence = in.Sequence
	wti.Witness = witness
	wti.SignatureScript = sigScript
	return wti
}

type TxVout struct {
	Scriptpubkey        string `json:"scriptpubkey"`
	ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
	ScriptpubkeyType    string `json:"scriptpubkey_type"`
	ScriptpubkeyAddress string `json:"scriptpubkey_address"`
	Value               int64  `json:"value"`
}

func (out *TxVout) TxOut() *wire.TxOut {
	pubSig, err := hex.DecodeString(out.Scriptpubkey)
	if err != nil {
		logx.Default.Errorf("err: %v", err)
		return nil
	}
	var wto = &wire.TxOut{}
	wto.Value = out.Value
	wto.PkScript = pubSig
	return wto
}

type BlockStatus struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight int64  `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int64  `json:"block_time"`
}
