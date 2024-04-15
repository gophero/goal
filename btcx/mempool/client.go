package mempool

import (
	"fmt"
	"github.com/gophero/goal/btcx"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

const (
	ApiUrl         = "https://mempool.space/api"
	ApiUrlTestNet3 = "https://mempool.space/testnet/api"
	ApiUrlSignet   = "https://mempool.space/signet/api"
)

type MempoolClient struct {
	baseURL string
	Net     btcx.Net
}

func NewClient(net btcx.Net) *MempoolClient {
	switch net {
	case btcx.MainNet:
		return &MempoolClient{baseURL: ApiUrl, Net: net}
	case btcx.TestNet:
		return &MempoolClient{baseURL: ApiUrlTestNet3, Net: net}
	case btcx.SigNet:
		return &MempoolClient{baseURL: ApiUrlTestNet3, Net: net}
	default:
		return nil
	}
}

func (c *MempoolClient) request(method, subPath string, requestBody io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, subPath)
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	return body, nil
}

func (c *MempoolClient) get(subPath string) ([]byte, error) {
	return c.request(http.MethodGet, subPath, nil)
}

func (c *MempoolClient) post(subPath string, body io.Reader) ([]byte, error) {
	return c.request(http.MethodPost, subPath, body)
}
