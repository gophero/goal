package twitter_test

import (
	"fmt"
	"github.com/gophero/goal/twitter"
	"testing"
)

func TestOAuth2TweetApi_RetweetBy(t *testing.T) {
	var tweetId = "1779812184761245967"
	users, _, err := twitter.OAuth2Apis.Tweet.RetweetBy(at, tweetId, nil)
	if err != nil {
		t.Errorf("test failed: %v", err)
	}
	fmt.Println(users)

	ff := twitter.NewFieldFilter()
	ff.AddUserField(twitter.UserFieldId, twitter.UserFieldProfileImageUrl, twitter.UserFieldCreatedAt, twitter.UserFieldVerified, twitter.UserFieldWithHeld, twitter.UserFieldDescription, twitter.UserFieldLocation)
	users, _, err = twitter.OAuth2Apis.Tweet.RetweetBy(at, tweetId, ff, twitter.GetParamOptions.MaxResults(1000), twitter.GetParamOptions.PaginationToken("test_token"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
}
