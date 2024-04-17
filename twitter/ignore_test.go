package twitter_test

import "github.com/gophero/goal/twitter"

var testEnv = struct {
	setting     twitter.Setting
	code        string
	accessToken string
	refrshToken string
	bearerToken string
}{
	setting: twitter.Setting{
		ClientId:     "YzdJSE1Db1FNenZ1bW8tZFlFSWw6MTpjaQ",
		ClientSecret: "s3Z00xyuWAaZ-lUq7OqWzdcNCJ5LzbchTWE-CY8z4xQiBOMPfi",
	},
	code:        "d3V5ZnJELVI1cElwUHJVMmJPaXBwNXUzdGhPMW0xVUh4WmFzbGdPR3VXSll1OjE3MTMyNTE1NDQyODE6MTowOmFjOjE",
	accessToken: "NHJGY0JrdllqZUNvLXJGXzRlSmZRZlJqMXk2U1Q3NUJJY1RzdjUtT0ZLUjdNOjE3MTMzNDkzMjE2NDU6MToxOmF0OjE",
	refrshToken: "ZkVIOVNKQWV0aHg0SGZGMmNHNkd6UWFPWU1OSEw0WGNvX2g1b0U0VGRBZTdpOjE3MTMzNDkzMjE2NDU6MToxOnJ0OjE",
	bearerToken: "AAAAAAAAAAAAAAAAAAAAAOm6tAEAAAAAd%2FFS0w4buJUaAHnmJpNOHk%2BkZ8A%3DxcdOBHtmhvyvWlM75lZBHuoC8Y0E7fm6cGSPaFKTeqO4stS4KG",
}
