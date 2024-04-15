package ord

import "github.com/pkg/errors"

const (
	mainNet = "mainnet"
	testNet = "testnet"
	signet  = "signet"
	regtest = "regtest"
)

type OrdSetting struct {
	Ord Setting `mapstructure:"ord" yaml:"ord"`
}

type Setting struct {
	BinPath       string `mapstructure:"binPath" yaml:"binPath"`
	ChainNet      string `mapstructure:"chainNet" yaml:"chainNet"`
	BitCoinRpcUrl string `mapstructure:"bitCoinRpcUrl" yaml:"bitCoinRpcUrl"`
	ServerUrl     string `mapstructure:"serverUrl" yaml:"serverUrl"`
}

func (s OrdSetting) checkNeccessary() error {
	if s.Ord.BinPath == "" {
		return errors.Wrap(InscribeError, "config error: need binPath")
	}
	if s.Ord.ChainNet == "" {
		return errors.Wrap(InscribeError, "config error: need chainNet")
	}
	if s.Ord.BitCoinRpcUrl == "" {
		return errors.Wrap(InscribeError, "config error: need bitCoinRpcUrl")
	}
	if s.Ord.ServerUrl == "" {
		return errors.Wrap(InscribeError, "config error: need serverUrl")
	}
	return nil
}
