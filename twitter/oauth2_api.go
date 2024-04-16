package twitter

import (
	"fmt"
)

const (
	oauth2ApiUrlFormat = "https://api.twitter.com/2%s"
)

var ApiError = fmt.Errorf("twitter api error")

var OAuth2Apis = NewOAuth2Api(NewOAuth2AuthApi(nil))

type OAuth2Api struct {
	Auth  *OAuth2AuthApi
	User  *OAuth2UserApi
	Tweet *OAuth2TweetApi
}

func NewOAuth2Api(authApi *OAuth2AuthApi) OAuth2Api {
	return OAuth2Api{
		Auth:  authApi,
		User:  NewOAuth2UserApi(),
		Tweet: NewOAuth2TweetApi(),
	}
}

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
}

type Result[T any] struct {
	Data T    `json:"data"`
	Meta Meta `json:"meta"`
}
