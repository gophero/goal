package twitter

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	userFieldsKey = "user.fields"
	expansionKey  = "expansions"
	tweetFieldKey = "tweet.fields"
)

type OAuth2UserApiFormParamOptions struct{}

type OAuth2UserApi struct{}

func NewOAuth2UserApi() *OAuth2UserApi {
	return &OAuth2UserApi{}
}

func (o *OAuth2UserApi) meUrl() string {
	return fmtUrl(oauth2ApiUrlFormat, "/users/me")
}

func (o *OAuth2UserApi) Me(accessToken string, ff *FieldFilter) (*UserInfo, error) {
	body := strings.NewReader(NewGetParam().FilterFields(ff).Param())
	req, err := http.NewRequest(http.MethodGet, o.meUrl(), body)
	if err != nil {
		return nil, errors.Wrapf(ApiError, "request error: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(ApiError, "request error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(ApiError, "request error: %v", resp.Status)
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Result[*UserInfo]
	if err := json.Unmarshal(bs, &result); err != nil {
		return nil, errors.Wrapf(ApiError, "invalid response: %v", string(bs))
	}
	return result.Data, nil
}

func (o *OAuth2UserApi) Followers(accessToken, id string, ff *FieldFilter, options ...GetParamOption) ([]*UserInfo, Meta, error) {
	url := fmtUrl(oauth2ApiUrlFormat, "/users/"+id+"/followers")
	params := NewGetParam().FilterFields(ff)
	for _, p := range options {
		p(params)
	}
	body := strings.NewReader(params.Param())
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

var EmptyUserInfo UserInfo

type UserInfo struct {
	Id                string        `json:"id"`
	Name              string        `json:"name"`
	Username          string        `json:"username"`
	CreatedAt         time.Time     `json:"created_at"`
	MostRecentTweetId string        `json:"most_recent_tweet_id"`
	Protected         bool          `json:"protected"`
	Withheld          any           `json:"withheld"`
	Location          string        `json:"location"`
	Url               string        `json:"url"`
	Description       string        `json:"description"`
	Verified          bool          `json:"verified"`
	Entities          Entities      `json:"entities"`
	ProfileImageUrl   string        `json:"profile_image_url"`
	PublicMetrics     PublicMetrics `json:"public_metrics"`
	PinnedTweetId     string        `json:"pinned_tweet_id"`
	Includes          []Include     `json:"includes"`
	Errors            Error         `json:"errors"`
}

type Withheld struct {
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type Entities struct {
	Url         []EntityUrl  `json:"url"`
	Description []EntityDesc `json:"description"`
}

type EntityUrl struct {
	Urls []EntityUrlItem `json:"urls"`
}

type EntityUrlItem struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Url         string `json:"url"`
	ExpandedUrl string `json:"expanded_url"`
	DisplayUrl  string `json:"display_url"`
}

type EntityDesc struct {
	EntityUrl
	Hashtags []EntityHashTag `json:"hashtags"`
	Mentions []EntityMention `json:"mentions"`
	Cashtags []EntityCashTag `json:"cashtags"`
}

type EntityHashTag struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Hashtag string `json:"hashtag"`
}

type EntityMention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Username string `json:"username"`
}

type EntityCashTag struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Cashtag string `json:"cashtag"`
}

type PublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type Include struct {
	Tweets []Tweet `json:"tweets"`
}
