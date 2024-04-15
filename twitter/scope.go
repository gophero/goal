package twitter

import "strings"

type Scope string

const (
	TweetRead          Scope = "tweet.read"           // All the Tweets you can view, including Tweets from protected accounts.
	TweetWrite               = "tweet.write"          // Tweet and Retweet for you.
	TweetModerateWrite       = "tweet.moderate.write" // Hide and unhide replies to your Tweets.
	UsersRead                = "users.read"           // Any account you can view, including protected accounts.
	FollowsRead              = "follows.read"         // People who follow you and people who you follow.
	FollowsWrite             = "follows.write"        // Follow and unfollow people for you.
	OfflineAccess            = "offline.access"       // Stay connected to your account until you revoke access.
	SpaceRead                = "space.read"           // All the Spaces you can view.
	MuteRead                 = "mute.read"            // Accounts you’ve muted.
	MuteWrite                = "mute.write"           // Mute and unmute accounts for you.
	LikeRead                 = "like.read"            // Tweets you’ve liked and likes you can view.
	LikeWrite                = "like.write"           // Like and un-like Tweets for you.
	ListRead                 = "list.read"            // Lists, list members, and list followers of lists you’ve created or are a member of, including private lists.
	ListWrite                = "list.write"           // Create and manage Lists for you.
	BlockRead                = "block.read"           // Accounts you’ve blocked.
	BlockWrite               = "block.write"          // Block and unblock accounts for you.
	BookmarkRead             = "bookmark.read"        // Get Bookmarked Tweets from an authenticated user.
	BookmarkWrite            = "bookmark.write"       // Bookmark and remove Bookmarks from Tweets
)

func formatScopes(scopes ...Scope) string {
	var sb strings.Builder
	for i, s := range scopes {
		sb.WriteString(string(s))
		if i < len(scopes)-1 {
			sb.WriteString("%20")
		}
	}
	return sb.String()
}
