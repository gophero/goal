package twitter

import (
	"encoding/json"

	"github.com/gophero/goal/errorx"
	"github.com/gophero/goal/logx"
)

type PostTweetParam struct {
	DirectMessageDeepLink string `json:"direct_message_deep_link"`
	ForSuperFollowersOnly bool   `json:"for_super_followers_only"`
	Geo                   GEO    `json:"geo"`
	Media                 Media  `json:"media"`
	Poll                  Poll   `json:"poll"`
	QuoteTweetId          string `json:"quote_tweet_id"`
	ReplySettings         string `json:"reply_settings"`
	Text                  string `json:"text"`
}

func (p *PostTweetParam) Json() []byte {
	if bs, err := json.Marshal(p); err != nil {
		logx.Default.Errorf("marshal json error: %v", errorx.Wrap(err))
		return []byte{}
	} else {
		return bs
	}
}

type GEO struct {
	PlaceId string `json:"place_id"`
}

type Media struct {
	MediaIds      []string `json:"media_ids"`
	TaggedUserIds []string `json:"tagged_user_ids"`
}

type Poll struct {
	DurationMinutes uint32   `json:"duration_minutes"`
	Options         []string `json:"options"`
}

type Reply struct {
	ExcludeReplyUserIds []string `json:"exclude_reply_user_ids"`
	InReplyToTweetId    string   `json:"in_reply_to_tweet_id"`
}

type PostTweetResp struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
