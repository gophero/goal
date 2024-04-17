package twitter_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gophero/goal/twitter"
)

var oauth2ApiUrlFormat = "https://api.twitter.com/2%s"

func TestOAuth2TweetApi_RetweetBy(t *testing.T) {
	var tweetId = "1779812184761245967"
	users, _, err := twitter.OAuth2Apis.Tweet.RetweetBy(testEnv.accessToken, tweetId, nil)
	if err != nil {
		t.Errorf("test failed: %v", err)
	}
	fmt.Println(users)

	ff := twitter.NewFieldFilter()
	ff.AddUserField(twitter.UserFieldId, twitter.UserFieldProfileImageUrl, twitter.UserFieldCreatedAt, twitter.UserFieldVerified, twitter.UserFieldWithHeld, twitter.UserFieldDescription, twitter.UserFieldLocation)
	users, _, err = twitter.OAuth2Apis.Tweet.RetweetBy(testEnv.accessToken, tweetId, ff, twitter.GetParamOptions.MaxResults(1000), twitter.GetParamOptions.PaginationToken("test_token"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
}

func TestPostTweet(t *testing.T) {
	var token = testEnv.accessToken
	url := fmt.Sprintf(oauth2ApiUrlFormat, "/tweets")
	var content = "{\"text\": \"Hello World!\"}"
	var body = strings.NewReader(content)
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	bs, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bs))
}

func TestPostTweetApi(t *testing.T) {
    var p twitter.PostTweetParam
    p.Text = "good morning!"
    r, err := twitter.OAuth2Apis.Tweet.PostTweet(testEnv.accessToken, p)
    if err != nil {
        t.Error(err)
    }
    fmt.Println(r)
}

func TestGetTweets(t *testing.T) {
	var token = testEnv.accessToken
	url := fmt.Sprintf(oauth2ApiUrlFormat, "/tweets")
    params := "1780259716532441535"
	var body = strings.NewReader(params)
	req, _ := http.NewRequest(http.MethodGet, url, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	bs, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bs))
}
