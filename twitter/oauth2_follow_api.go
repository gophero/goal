package twitter

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Oauth2FollowApi struct{}

func (o *Oauth2FollowApi) followUrl(userId string) string {
	return fmtUrl(followUrl, userId)
}

func (o *Oauth2FollowApi) Follow(accessToken string, userId, targetUserId string) (FollowRet, error) {
	url := o.followUrl(userId)
	body := strings.NewReader(
		NewGetParam().
			Append("target_user_id", targetUserId).
			Param(),
	)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return FollowRet{}, errors.Wrapf(ApiError, "request error: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return FollowRet{}, errors.Wrapf(ApiError, "request error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return FollowRet{}, errors.Wrapf(ApiError, "request error: %v", resp.Status)
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return FollowRet{}, err
	}
	var result Result[FollowRet]
	if err := json.Unmarshal(bs, &result); err != nil {
		return FollowRet{}, errors.Wrapf(ApiError, "invalid response: %v", string(bs))
	}
	return result.Data, nil
}

func (o *Oauth2FollowApi) UnFollow() {
}

type FollowRet struct {
	Following     bool `json:"following"`
	PendingFollow bool `json:"pending_follow"`
}
