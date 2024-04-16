package twitter_test

import (
	"encoding/json"
	"fmt"
	"github.com/gophero/goal/twitter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMe(t *testing.T) {
	user, err := twitter.OAuth2Apis.User.Me(at, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
	assert.True(t, user.Id != "" && user.Name != "" && user.Username != "" && user.ProfileImageUrl == "")
	ff := twitter.NewFieldFilter()
	ff.AddUserField(twitter.UserFieldId, twitter.UserFieldProfileImageUrl, twitter.UserFieldCreatedAt, twitter.UserFieldVerified, twitter.UserFieldWithHeld, twitter.UserFieldDescription, twitter.UserFieldLocation)
	user, err = twitter.OAuth2Apis.User.Me(at, ff)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
	assert.True(t, user.CreatedAt.Year() > 2000 && user.ProfileImageUrl != "")
}

func TestFollowers(t *testing.T) {
	id := "1776491059318792192"
	users, err := twitter.OAuth2Apis.User.Followers(at, id, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
	ff := twitter.NewFieldFilter()
	ff.AddUserField(twitter.UserFieldId, twitter.UserFieldProfileImageUrl, twitter.UserFieldCreatedAt, twitter.UserFieldVerified, twitter.UserFieldWithHeld, twitter.UserFieldDescription, twitter.UserFieldLocation)
	users, err = twitter.OAuth2Apis.User.Followers(at, id, ff, twitter.OAuth2Apis.User.Param.MaxResults(1000), twitter.OAuth2Apis.User.Param.PaginationToken("test_token"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
}

func TestModel(t *testing.T) {
	js := `
{
  "data": [
    {
      "id": "6253282",
      "name": "Twitter API",
      "username": "TwitterAPI"
    },
    {
      "id": "2244994945",
      "name": "Twitter Dev",
      "username": "TwitterDev"
    },
    {
      "id": "783214",
      "name": "Twitter",
      "username": "Twitter"
    },
    {
      "id": "95731075",
      "name": "Twitter Safety",
      "username": "TwitterSafety"
    },
    {
      "id": "3260518932",
      "name": "Twitter Moments",
      "username": "TwitterMoments"
    },
    {
      "id": "373471064",
      "name": "Twitter Music",
      "username": "TwitterMusic"
    },
    {
      "id": "791978718",
      "name": "Twitter Official Partner",
      "username": "OfficialPartner"
    },
    {
      "id": "17874544",
      "name": "Twitter Support",
      "username": "TwitterSupport"
    },
    {
      "id": "234489024",
      "name": "Twitter Comms",
      "username": "TwitterComms"
    },
    {
      "id": "1526228120",
      "name": "Twitter Data",
      "username": "TwitterData"
    }
  ],
  "meta": {
    "result_count": 10,
    "next_token": "DFEDBNRFT3MHCZZZ"
  }
}
`
	var r twitter.Result[[]*twitter.UserInfo]
	if err := json.Unmarshal([]byte(js), &r); err != nil {
		t.Errorf("test failed: %v", err)
	} else {
		users := r.Data
		assert.True(t, r.Meta.ResultCount == 10)
		assert.True(t, uint32(len(users)) == r.Meta.ResultCount)
		assert.True(t, r.Meta.NextToken == "DFEDBNRFT3MHCZZZ")
	}
}
