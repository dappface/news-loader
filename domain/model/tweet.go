package model

import (
	"encoding/json"
	"time"

	"github.com/tidwall/gjson"
)

type Tweet struct {
	CreatedAt time.Time `firestore:"createdAt"            json:"-"`
	// [Unused] ID        int64     `firestore:"id"                   json:"id"`
	IDStr     string `firestore:"idStr"                json:"id_str"`
	Text      string `firestore:"text"                 json:"text"`
	Source    string `firestore:"source"               json:"source"`
	Truncated bool   `firestore:"truncated"            json:"truncated"`
	// [Unused] InReplyToStatusID    *int64       `firestore:"inReplyToStatusId"    json:"in_reply_to_status_id"`
	InReplyToStatusIDStr *string `firestore:"inReplyToStatusIdStr" json:"in_reply_to_status_id_str"`
	// [Unused] InReplyToUserID      *int64       `firestore:"inReplyToUserId"      json:"in_reply_to_user_id"`
	InReplyToUserIDStr  *string      `firestore:"inReplyToUserIdStr"   json:"in_reply_to_user_id_str"`
	InReplyToScreenName *string      `firestore:"inReplyToScreenName"  json:"in_reply_to_screen_name"`
	User                User         `firestore:"user"                 json:"user,User"`
	Coordinates         *Coordinates `firestore:"coordinates"          json:"coordinates"`
	Place               *Place       `firestore:"place"                json:"place"`
	// [Unused] QuotedStatusID       *int64       `firestore:"quotedStatusId"       json:"quoted_status_id"`
	QuotedStatusIDStr *string `firestore:"quotedStatusIdStr"    json:"quoted_status_id_str"`
	IsQuoteStatus     bool    `firestore:"isQuoteStatus"        json:"is_quote_status"`
	QuotedStatus      *Tweet  `firestore:"quotedStatus"         json:"quoted_status"`
	RetweetedStatus   *Tweet  `firestore:"retweetedStatus"      json:"retweeted_status"`
	// [Unused] QuoteCount           *int
	ReplyCount       int               `firestore:"replyCount"       json:"reply_count"`
	RetweetCount     int               `firestore:"retweetCount"     json:"retweet_count"`
	FavoriteCount    *int              `firestore:"favoriteCount"    json:"favorite_count"`
	Entities         Entities          `firestore:"entities"         json:"entities"`
	ExtendedEntities *ExtendedEntities `firestore:"extendedEntities" json:"extended_entities"`
	Favorited        *bool             `firestore:"favorited"        json:"favorited"`
	Retweeted        bool              `firestore:"retweeted"        json:"retweeted"`
	// [Unused] PossiblySensitive *bool
	// [Unused] FilterLevel string
	Lang *string `firestore:"lang" json:"lang"`
	// [Unused] MatchingRules []Rule
}

func (t *Tweet) UnmarshalJSON(data []byte) error {
	type t2 Tweet
	if err := json.Unmarshal(data, (*t2)(t)); err != nil {
		return err
	}

	jsonS := string(data)
	createdAt := gjson.Get(jsonS, "created_at").String()
	tm, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", createdAt)
	if err != nil {
		return err
	}
	t.CreatedAt = tm
	return nil
}

type LatestTweet struct {
	ID        string    `firestore:"id"`
	ListName  string    `firestore:"listName"`
	CreatedAt time.Time `firestore:"createdAt"`
	UpdatedAt time.Time `firestore:"updatedAt"`
}
