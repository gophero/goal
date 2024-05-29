package twitter

import (
	"bytes"
	"encoding/json"
	"github.com/gophero/goal/errorx"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type OAuth2FollowApi struct{}

func NewOAuth2FollowApi() *OAuth2FollowApi {
	return &OAuth2FollowApi{}
}

func (o *OAuth2FollowApi) followUrl(userId string) string {
	return fmtUrl(followUrl, userId)
}

func (o *OAuth2FollowApi) Follow(accessToken string, userId, targetUserId string) (FollowRet, error) {
	url := o.followUrl(userId)
	m := map[string]any{}
	m["target_user_id"] = targetUserId
	bs, err := json.Marshal(m)
	if err != nil {
		return FollowRet{}, errorx.Wrapf(err, "param error")
	}
	body := bytes.NewReader(bs)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return FollowRet{}, errors.Wrapf(ApiError, "request error: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return FollowRet{}, errors.Wrapf(ApiError, "request error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return FollowRet{}, errors.Wrapf(ApiError, "request error: %v", resp.Status)
	}

	bs, err = io.ReadAll(resp.Body)
	if err != nil {
		return FollowRet{}, err
	}
	var result Result[FollowRet]
	if err := json.Unmarshal(bs, &result); err != nil {
		return FollowRet{}, errors.Wrapf(ApiError, "invalid response: %v", string(bs))
	}
	return result.Data, nil
}

func (o *OAuth2FollowApi) UnFollow() {
}

type FollowRet struct {
	Following     bool `json:"following"`
	PendingFollow bool `json:"pending_follow"`
}
