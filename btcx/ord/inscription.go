package ord

type Inscription struct {
	ID        string `json:"id"`
	Location  string `json:"location"`
	Explorler string `json:"explorer"`
	Postage   int64  `json:"postage"`
}

type InscribeResult struct {
	Commit       string         `json:"commit"`
	CommitPsbt   interface{}    `json:"commit_psbt"`
	Inscriptions []*Inscription `json:"inscriptions"`
	Parent       string         `json:"parent"`
	Reveal       string         `json:"reveal"`
	RevealPsbt   interface{}    `json:"reveal_psbt"`
	TotalFees    int64          `json:"total_fees"`
}
