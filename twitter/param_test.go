package twitter_test

import (
	"fmt"
	"github.com/gophero/goal/twitter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParam(t *testing.T) {
	ff := twitter.NewFieldFilter()
	ff.AddUserField(twitter.UserFieldId, twitter.UserFieldProfileImageUrl, twitter.UserFieldCreatedAt, twitter.UserFieldVerified, twitter.UserFieldWithHeld, twitter.UserFieldDescription, twitter.UserFieldLocation)
	getParams := twitter.NewGetParam()
	getParams.FilterFields(ff)
	fmt.Println(getParams.Param())
	assert.True(t, getParams.Param() == "user.fields=id,profile_image_url,created_at,verified,withheld,description,location")
}
