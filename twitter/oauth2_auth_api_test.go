package twitter_test

import (
	"fmt"
	"github.com/gophero/goal/twitter"
	"testing"
)

var testSetting = twitter.Setting{
	ApiKey:            "XEKRjx8NXBGvukzo5XHzamvu1",
	ApiSecret:         "gh8k0LrxYTu60XvQRpDFw0Ssrph2c0d9C82eurRri8zS2ppt1b",
	BearerToken:       "AAAAAAAAAAAAAAAAAAAAABfLsgEAAAAAOVvVQ1ervFulfzToTDOIqvrv9CI%3DxuKGaRmqu0Y5dwd0674nJIvUYotwkIvI8NucNoFIpL9HpcI0UT",
	AccessToken:       "1320117828-HrjSjn2vYTnwt3ZZjvB554XHkBLz9V9ezdo2Wxy",
	AccessTokenSecret: "hQeQmSAAMNbCSy46WCIRh9IXqnFXHTKXu0DBE4TZSKAZt",
	ClientId:          "Y0g0SjdTME1acW9iMFMwTEJyZ186MTpjaQ",
	ClientSecret:      "LbPfZ1TwJSNtBYCIbyESJcfhEhk6ihto94imb3UY2wwr4ggX_I",
}

var scopes = "offline.access%20tweet.read%20users.read%20follows.read%20follows.write"
var url = "https://twitter.com/i/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=state&code_challenge=%s&code_challenge_method=plain"

func TestRequestToken(t *testing.T) {
	var code = ""
	var state = ""
	var redirectUri = ""
	ak, err := twitter.OAuth2Apis.AuthApi.RequestAccessToken(testSetting.ClientId, testSetting.ClientSecret, code, state, redirectUri)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ak)
}

func TestRefreshToken(t *testing.T) {
	refreshToken := "Sm8zQTNhZWFVODJvZ1MwUG5PZzlicl9tSm9NV2ktRklQWERrb1FEUms2empkOjE3MTIzNzU3MzA5ODY6MToxOmF0OjE"
	ss, err := twitter.OAuth2Apis.UserApi.Me(refreshToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ss)
}

func TestRevokeToken(t *testing.T) {
	token := "N2IwS2J6dmpZeF95enFYSUFSQnduSkhYb1VvTTZnMEhnOUNtUkRLMVJpcnBJOjE3MTE1MzM1NTIyNzg6MTowOmF0OjE"
	err := twitter.OAuth2Apis.AuthApi.RevokeAccessToken(testSetting.ClientId, testSetting.ClientSecret, token)
	if err != nil {
		t.Fatal(err)
	}
}
