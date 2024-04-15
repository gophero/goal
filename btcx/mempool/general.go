package mempool

import (
	"encoding/json"
	"github.com/gophero/goal/btcx/mempool/model"
)

func (c *MempoolClient) GetPrices() (model.Price, error) {
	if bs, err := c.get("/v1/prices"); err != nil {
		return model.Price{}, nil
	} else {
		var p model.Price
		return p, json.Unmarshal(bs, &p)
	}
}

func (c *MempoolClient) GetDifficultyAdjustment() (model.DifficultyAdjust, error) {
	if bs, err := c.get("/v1/difficulty-adjustment"); err != nil {
		return model.DifficultyAdjust{}, err
	} else {
		var r model.DifficultyAdjust
		return r, json.Unmarshal(bs, &r)
	}
}
