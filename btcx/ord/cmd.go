package ord

import (
	"fmt"
	"strconv"
	"strings"
)

var Cmd command

const (
	paramChain         = "--chain"
	paramBitcoinRpcUrl = "--bitcoin-rpc-url"
	paramWalletName    = "--name"
	paramServerUrl     = "--server-url"
	paramBatch         = "--batch"
	paramCborMetadata  = "--cbor-metadata"
	paramCommitFeeRate = "--commit-fee-rate"
	paramCompress      = "--compress"
	paramDelegate      = "--delegate"
	paramDestination   = "--destination"
	paramDryRun        = "--dry-run"
	paramFeeRate       = "--fee-rate"
	paramFile          = "--file"
	paramJsonMetadata  = "--json-metadata"
	paramMetaProtocol  = "--metaprotocol"
	paramNoBackup      = "--no-backup"
	paramNoLimit       = "--no-limit"
	paramParent        = "--parent"
	paramPostage       = "--postage"
	paramReinscribe    = "--reinscribe"
	paramSat           = "--sat"
	paramSatpoint      = "--satpoint"

	baseCmd         = "%s --chain %s --bitcoin-rpc-url %s"
	space           = " "
	cmdWallet       = "wallet"
	cmdInscriptions = "inscriptions"
	cmdInscribe     = "inscribe"
)

type command struct{}

type builder struct {
	w strings.Builder
}

func newBuilder() *builder {
	return &builder{w: strings.Builder{}}
}

func (b *builder) write(s string) *builder {
	b.w.WriteString(s)
	return b
}

func (b *builder) space() *builder {
	b.w.WriteString(space)
	return b
}

func (b *builder) equal() *builder {
	b.w.WriteString("=")
	return b
}

func (b *builder) String() string {
	return b.w.String()
}

func (c command) inscribeCmd(setting OrdSetting, req *InscribeRequest) (string, error) {
	if err := setting.checkNeccessary(); err != nil {
		return "", err
	}

	cmd := fmt.Sprintf(baseCmd, setting.Ord.BinPath, setting.Ord.ChainNet, setting.Ord.BitCoinRpcUrl)
	feeRateStr := strconv.Itoa(int(req.FeeRate))
	var builder = newBuilder()
	builder.write(cmd)
	builder.space().write(cmdWallet)
	builder.space().write(paramWalletName).equal().write(req.Wallet)
	builder.space().write(paramServerUrl).space().write(setting.Ord.ServerUrl)
	builder.space().write(cmdInscribe)
	builder.space().write(paramFeeRate).space().write(feeRateStr)
	if req.File != "" {
		builder.space().write(paramFile).space().write(req.File)
	}
	if req.BatchFile != "" {
		builder.space().write(paramBatch).space().write(req.BatchFile)
	}
	if req.Postage > 0 {
		postageStr := strconv.Itoa(int(req.Postage)) + "sat"
		builder.space().write(paramPostage).space().write(postageStr)
	}
	if req.ParentInscriptionId != "" {
		builder.space().write(paramParent).space().write(req.ParentInscriptionId)
	}
	if req.DelegateInscriptionId != "" {
		builder.space().write(paramDelegate).space().write(req.DelegateInscriptionId)
	}
	if req.DryRun {
		builder.space().write(paramDryRun)
	}
	return builder.String(), nil
}

func (c command) queryInscriptionCmd(setting OrdSetting, req *InscriptionQueryRequest) (string, error) {
	if err := setting.checkNeccessary(); err != nil {
		return "", err
	}

	cmds := fmt.Sprintf(baseCmd, setting.Ord.BinPath, setting.Ord.ChainNet, setting.Ord.BitCoinRpcUrl)
	var builder = newBuilder()
	builder.write(cmds)
	builder.space().write(cmdWallet)
	builder.space().write(paramWalletName).equal().write(req.Wallet)
	builder.space().write(paramServerUrl).space().write(setting.Ord.ServerUrl)
	builder.space().write(cmdInscriptions)
	return builder.String(), nil
}
