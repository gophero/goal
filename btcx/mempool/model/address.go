package model

type AddressInfo struct {
	Address      string       `json:"address"`
	ChainStats   ChainStats   `json:"chain_stats"`
	MempoolStats MempoolStats `json:"mempool_stats"`
}

type ValidAddress struct {
	Isvalid        bool   `json:"isvalid"`
	Address        string `json:"address"`
	ScriptPubKey   string `json:"scriptPubKey"`
	Isscript       bool   `json:"isscript"`
	Iswitness      bool   `json:"iswitness"`
	WitnessVersion int32  `json:"witness_version"`
	WitnessProgram string `json:"witness_program"`
}

type ChainStats struct {
	FundedTxoCount int64 `json:"funded_txo_count"`
	FundedTxoSum   int64 `json:"funded_txo_sum"`
	SpentTxoCount  int64 `json:"spent_txo_count"`
	SpentTxoSum    int64 `json:"spent_txo_sum"`
	TxCount        int64 `json:"tx_count"`
}

type MempoolStats struct {
	FundedTxoCount int64 `json:"funded_txo_count"`
	FundedTxoSum   int64 `json:"funded_txo_sum"`
	SpentTxoCount  int64 `json:"spent_txo_count"`
	SpentTxoSum    int64 `json:"spent_txo_sum"`
	TxCount        int64 `json:"tx_count"`
}
