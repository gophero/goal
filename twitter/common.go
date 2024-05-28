package twitter

import (
	"fmt"
	"strings"

	"github.com/gophero/goal/collection/slicex"
)

var ApiError = fmt.Errorf("twitter api error")

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
}

type Result[T any] struct {
	Data T    `json:"data"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	ResultCount   uint32 `json:"result_count"`
	PreviousToken string `json:"previous_token"`
	NextToken     string `json:"next_token"`
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

type (
	Expansion    string
	TwitterField string
	UserField    string
)

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
