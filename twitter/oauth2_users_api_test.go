package twitter_test

import (
	"fmt"
	"github.com/gophero/goal/twitter"
	"testing"
)

func TestMe(t *testing.T) {
	at := "ak5PUU16STFieGphUURjRloyQUZkLUtKTWpvX1ZNb1E2YkFwUXVHOW1hY3JnOjE3MTIzNzY1MDQ2MDY6MTowOmF0OjE"
	user, err := twitter.OAuth2Apis.UserApi.Me(at, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
	ff := twitter.NewFieldFilter()
	ff.AddUserField(twitter.UserFieldId, twitter.UserFieldProfileImageUrl, twitter.UserFieldCreatedAt, twitter.UserFieldVerified, twitter.UserFieldWithHeld, twitter.UserFieldDescription, twitter.UserFieldLocation)
	user, err = twitter.OAuth2Apis.UserApi.Me(at, ff)
	if err != nil {
		t.Fatal(err)
	}
	// {"data":{"name":"Cuplis men","profile_image_url":"https://pbs.twimg.com/profile_images/1587722231576113152/LQCzO8Lf_normal.png","username":"NetBit_X","created_at":"2013-04-01T11:16:55.000Z","verified":false,"id":"1320117828","description":"Kukan slalu buat yg ter baik"}}
	fmt.Println(user)
}
