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
	ak, err := twitter.OAuth2Apis.Auth.RequestAccessToken(testEnv.setting.ClientId, testEnv.setting.ClientSecret, testEnv.code, state, redirectUri)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ak)
}

func TestRefreshToken(t *testing.T) {
	refreshToken := testEnv.refrshToken
	ss, err := twitter.OAuth2Apis.Auth.RefreshAccessToken(testEnv.setting.ClientId, testEnv.setting.ClientSecret, refreshToken)
	if err != nil {
		t.Fatal(err)
	}
	testEnv.accessToken = ss.AccessToken
	testEnv.refrshToken = ss.RefreshToken
}

func TestRevokeToken(t *testing.T) {
	token := testEnv.accessToken
	err := twitter.OAuth2Apis.Auth.RevokeAccessToken(testEnv.setting.ClientId, testEnv.setting.ClientSecret, token)
	if err != nil {
		t.Fatal(err)
	}
}
