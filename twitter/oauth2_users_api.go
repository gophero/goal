package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/gophero/goal/collection/slicex"
	"github.com/gophero/goal/stringx"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	userFieldsKey = "user.fields"
	expansionKey  = "expansions"
	tweetFieldKey = "tweet.fields"
)

type OAuth2UserApi struct {
}

func NewOAuth2UserApi() OAuth2AuthApi {
	return OAuth2AuthApi{}
}

func (o OAuth2UserApi) meUrl() string {
	return fmt.Sprintf(oauth2ApiUrlFormat, "/users/me")
}

type FieldFilter struct {
	Expansions    []Expansion
	TwitterFields []TwitterField
	UserFields    []UserField
}

func NewFieldFilter() *FieldFilter {
	return &FieldFilter{}
}

func (ff *FieldFilter) AddExpansion(exps ...Expansion) *FieldFilter {
	ff.Expansions = append(ff.Expansions, exps...)
	return ff
}

func (ff *FieldFilter) AddTwitterField(tfs ...TwitterField) *FieldFilter {
	ff.TwitterFields = append(ff.TwitterFields, tfs...)
	return ff
}

func (ff *FieldFilter) AddUserField(ufs ...UserField) *FieldFilter {
	ff.UserFields = append(ff.UserFields, ufs...)
	return ff
}

func (o OAuth2UserApi) Me(accessToken string, ff *FieldFilter) (UserInfo, error) {
	var params = ""
	if ff != nil {
		var builder = stringx.NewBuilder()
		if len(ff.Expansions) > 0 {
			builder.WriteString("expansions=").WriteString(formatExpansion(ff.Expansions...)).WriteString("&")
		}
		if len(ff.UserFields) > 0 {
			builder.WriteString("user.fields=").WriteString(formatUserFields(ff.UserFields...)).WriteString("&")
		}
		if len(ff.TwitterFields) > 0 {
			builder.WriteString("tweet.fields=").WriteString(formatTweetFields(ff.TwitterFields...))
		}
		params = strings.TrimRight(builder.String(), "&")
	}
	var body = strings.NewReader(params)
	req, err := http.NewRequest(http.MethodGet, o.meUrl(), body)
	if err != nil {
		return EmptyUserInfo, errors.Wrapf(ApiError, "request error: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return EmptyUserInfo, errors.Wrapf(ApiError, "request error: %v", err)
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return EmptyUserInfo, err
	}
	var result Result[UserInfo]
	if err := json.Unmarshal(bs, &result); err != nil {
		return EmptyUserInfo, errors.Wrapf(ApiError, "invalid response: %v", string(bs))
	}
	return result.Data, nil
}

type Expansion string
type TwitterField string
type UserField string

const (
	ExpansionPinnedTweetId Expansion = "pinned_tweet_id"
)

const (
	TwitterFieldAttachments        TwitterField = "attachments"
	TwitterFieldAuthorId                        = "author_id"
	TwitterFieldContextAnnotations              = "context_annotations"
	TwitterFieldConversationId                  = "conversation_id"
	TwitterFieldCreatedAt                       = "created_at"
	TwitterFieldEditControls                    = "edit_controls"
	TwitterFieldEntities                        = "entities"
	TwitterFieldGeo                             = "geo"
	TwitterFieldId                              = "id"
	TwitterFieldInReplyToUserId                 = "in_reply_to_user_id"
	TwitterFieldLang                            = "lang"
	TwitterFieldNonPublicMetrics                = "non_public_metrics"
	TwitterFieldPublicMetrics                   = "public_metrics"
	TwitterFieldOrganicMetrics                  = "organic_metrics"
	TwitterFieldPromotedMetrics                 = "promoted_metrics"
	TwitterFieldPossiblySensitive               = "possibly_sensitive"
	TwitterFieldReferencedTweets                = "referenced_tweets"
	TwitterFieldReplySettings                   = "reply_settings"
	TwitterFieldSource                          = "source "
	TwitterFieldText                            = "text"
	TwitterFieldWithheld                        = "withheld"
)

const (
	UserFieldCreatedAt         UserField = "created_at"
	UserFieldDescription                 = "description"
	UserFieldEntities                    = "entities"
	UserFieldId                          = "id"
	UserFieldLocation                    = "location"
	UserFieldMostRecentTweetId           = "most_recent_tweet_id"
	UserFieldName                        = "name"
	UserFieldPinnedTweetId               = "pinned_tweet_id"
	UserFieldProfileImageUrl             = "profile_image_url"
	UserFieldProtected                   = "protected"
	UserFieldPublicMetrics               = "public_metrics"
	UserFieldUrl                         = "url"
	UserFieldUserName                    = "username"
	UserFieldVerified                    = "verified"
	UserFieldVerifiedType                = "verified_type"
	UserFieldWithHeld                    = "withheld"
)

func formatExpansion(exps ...Expansion) string {
	rs := slicex.Eachv(exps, func(v Expansion) string {
		return string(v)
	})
	return strings.Join(rs, ",")
}

func formatTweetFields(tfs ...TwitterField) string {
	rs := slicex.Eachv(tfs, func(v TwitterField) string {
		return string(v)
	})
	return strings.Join(rs, ",")
}

func formatUserFields(ufs ...UserField) string {
	rs := slicex.Eachv(ufs, func(v UserField) string {
		return string(v)
	})
	return strings.Join(rs, ",")
}

var EmptyUserInfo UserInfo

type UserInfo struct {
	Id                string        `json:"id"`
	Name              string        `json:"name"`
	Username          string        `json:"username"`
	CreatedAt         time.Time     `json:"created_at"`
	MostRecentTweetId string        `json:"most_recent_tweet_id"`
	Protected         bool          `json:"protected"`
	Withheld          any           `json:"withheld"`
	Location          string        `json:"location"`
	Url               string        `json:"url"`
	Description       string        `json:"description"`
	Verified          bool          `json:"verified"`
	Entities          Entities      `json:"entities"`
	ProfileImageUrl   string        `json:"profile_image_url"`
	PublicMetrics     PublicMetrics `json:"public_metrics"`
	PinnedTweetId     string        `json:"pinned_tweet_id"`
	Includes          []Include     `json:"includes"`
	Errors            Error         `json:"errors"`
}

type Withheld struct {
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type Entities struct {
	Url         []EntityUrl  `json:"url"`
	Description []EntityDesc `json:"description"`
}

type EntityUrl struct {
	Urls []EntityUrlItem `json:"urls"`
}

type EntityUrlItem struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Url         string `json:"url"`
	ExpandedUrl string `json:"expanded_url"`
	DisplayUrl  string `json:"display_url"`
}

type EntityDesc struct {
	EntityUrl
	Hashtags []EntityHashTag `json:"hashtags"`
	Mentions []EntityMention `json:"mentions"`
	Cashtags []EntityCashTag `json:"cashtags"`
}

type EntityHashTag struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Hashtag string `json:"hashtag"`
}

type EntityMention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Username string `json:"username"`
}

type EntityCashTag struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Cashtag string `json:"cashtag"`
}

type PublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type Include struct {
	Tweets []Tweet `json:"tweets"`
}
