package model

type DifficultyAdjust struct {
	ProgressPercent       float64 `json:"progressPercent"`
	DifficultyChange      float64 `json:"difficultyChange"`
	EstimatedRetargetDate int64   `json:"estimatedRetargetDate"`
	RemainingBlocks       int64   `json:"remainingBlocks"`
	RemainingTime         int64   `json:"remainingTime"`
	PreviousRetarget      float64 `json:"previousRetarget"`
	NextRetargetHeight    int64   `json:"nextRetargetHeight"`
	TimeAvg               int64   `json:"timeAvg"`
	AdjustedTimeAvg       int64   `json:"adjustedTimeAvg"`
	TimeOffset            int64   `json:"timeOffset"`
}

type Price struct {
	Time int64 `json:"time"`
	USD  int64 `json:"USD"`
	EUR  int64 `json:"EUR"`
	GBP  int64 `json:"GBP"`
	CAD  int64 `json:"CAD"`
	CHF  int64 `json:"CHF"`
	AUD  int64 `json:"AUD"`
	JPY  int64 `json:"JPY"`
}
