package btcx

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

type Net string

const (
	MainNet Net = "mainnet"
	TestNet Net = "testnet3"
	SigNet  Net = "signet"
	Regtest Net = "regtest"
)

func (n Net) Params() *chaincfg.Params {
	switch n {
	case MainNet:
		return &chaincfg.MainNetParams
	case TestNet:
		return &chaincfg.TestNet3Params
	case SigNet:
		return &chaincfg.SigNetParams
	case Regtest:
		return &chaincfg.SimNetParams // TODO not sure it is correct
	default:
		return nil
	}
}

type UnspentOutput struct {
	Outpoint *wire.OutPoint
	Output   *wire.TxOut
}

type Client interface {
	GetRawTransaction(txHash *chainhash.Hash) (*wire.MsgTx, error)

	BroadcastTx(tx *wire.MsgTx) (*chainhash.Hash, error)

	ListUnspent(address btcutil.Address) ([]*UnspentOutput, error)
}

type RpcClient struct {
}

func (r *RpcClient) GetRawTransaction(txHash *chainhash.Hash) (*wire.MsgTx, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RpcClient) BroadcastTx(tx *wire.MsgTx) (*chainhash.Hash, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RpcClient) ListUnspent(address btcutil.Address) ([]*UnspentOutput, error) {
	panic("not implemented") // TODO: Implement
}
