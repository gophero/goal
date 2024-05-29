package twitter_test

import (
	"fmt"
	"os"

	"github.com/gophero/goal/twitter"
)

var testEnv testenv

type testenv struct {
	setting     twitter.Setting
	code        string
	accessToken string
	refrshToken string
	bearerToken string
}

func init() {
	testEnv.setting = twitter.Setting{
		ClientId:     os.Getenv("x_clientid"),
		ClientSecret: os.Getenv("x_clientsecret"),
	}
	testEnv.code = ""
	testEnv.accessToken = "Z1JsZTRlZDhiMXJxZU0zS3hhTWt5dFZ6S3h2UzhDS0NBNUZQV3J4YkZ0dkQwOjE3MTY5NTA2ODA3NDE6MToxOmF0OjE" //os.Getenv("x_accesstoken")
	testEnv.refrshToken = os.Getenv("x_refreshtoken")
	testEnv.bearerToken = os.Getenv("x_bearertoken")
	fmt.Println("testenv:", testEnv)
}
