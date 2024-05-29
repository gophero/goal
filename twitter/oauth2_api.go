package twitter

var OAuth2Apis = NewOAuth2Api(NewOAuth2AuthApi(nil))

type OAuth2Api struct {
	Auth   *OAuth2AuthApi
	User   *OAuth2UserApi
	Tweet  *OAuth2TweetApi
	Follow *OAuth2FollowApi
}

func NewOAuth2Api(authApi *OAuth2AuthApi) *OAuth2Api {
	return &OAuth2Api{
		Auth:   authApi,
		User:   NewOAuth2UserApi(),
		Tweet:  NewOAuth2TweetApi(),
		Follow: NewOAuth2FollowApi(),
	}
}
