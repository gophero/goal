package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

type OAuth2TweetApi struct {
}

func NewOAuth2TweetApi() *OAuth2TweetApi {
	return &OAuth2TweetApi{}
}

func (o *OAuth2TweetApi) RetweetBy(accessToken, tweetId string, ff *FieldFilter, options ...GetParamOption) ([]*UserInfo, Meta, error) {
	var url = fmt.Sprintf(oauth2ApiUrlFormat, "/tweets/"+tweetId+"/retweeted_by")
	var params = NewGetParam().FilterFields(ff)
	for _, p := range options {
		p(params)
	}
	var body = strings.NewReader(params.Param())
	req, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return []*UserInfo{}, Meta{}, errors.Wrapf(ApiError, "request error: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []*UserInfo{}, Meta{}, errors.Wrapf(ApiError, "request error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return []*UserInfo{}, Meta{}, errors.Wrapf(ApiError, "request error: %v", resp.Status)
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return []*UserInfo{}, Meta{}, err
	}
	var result Result[[]*UserInfo]
	if err := json.Unmarshal(bs, &result); err != nil {
		return []*UserInfo{}, Meta{}, errors.Wrapf(ApiError, "invalid response: %v", string(bs))
	}
	return result.Data, result.Meta, nil
}

type Tweet struct {
}
