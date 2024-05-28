package twitter

import "fmt"

const (
	// auth
	oauth2ApiUrlFormat      = "https://api.twitter.com/2%s"
	auth2AuthorizeUrlFormat = "https://twitter.com/i/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s&code_challenge=%s&code_challenge_method=plain"
	// follow
	followUrl = "https://api.twitter.com/2/users/%s/following"
)

func fmtUrl(url string, ps ...any) string {
	return fmt.Sprintf(url, ps...)
}
