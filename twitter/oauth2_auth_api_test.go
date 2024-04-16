package twitter_test

import (
	"fmt"
	"github.com/gophero/goal/twitter"
	"testing"
)

var scopes = "offline.access%20tweet.read%20users.read%20follows.read%20follows.write"
var url = "https://twitter.com/i/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=state&code_challenge=%s&code_challenge_method=plain"

func TestRequestToken(t *testing.T) {
	var state = ""
	var redirectUri = "http://localhost:8080"
	ak, err := twitter.OAuth2Apis.Auth.RequestAccessToken(testSetting.ClientId, testSetting.ClientSecret, code, state, redirectUri)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ak)
}

func TestRefreshToken(t *testing.T) {
	refreshToken := rt
	ss, err := twitter.OAuth2Apis.Auth.RefreshAccessToken(testSetting.ClientId, testSetting.ClientSecret, refreshToken)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ss)
}

func TestRevokeToken(t *testing.T) {
	token := at
	err := twitter.OAuth2Apis.Auth.RevokeAccessToken(testSetting.ClientId, testSetting.ClientSecret, token)
	if err != nil {
		t.Fatal(err)
	}
}
