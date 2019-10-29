package model

import "time"

type Entities struct {
	Hashtags     []Hashtag     `firestore:"hashtags"     json:"hashtags"`
	Media        *[]Media      `firestore:"media"        json:"media"`
	Symbols      []Symbol      `firestore:"symbols"      json:"symbols"`
	URLs         []URL         `firestore:"urls"         json:"urls"`
	UserMentions []UserMention `firestore:"userMentions" json:"user_mentions"`
	Polls        *[]Poll       `firestore:"polls"        json:"polls"`
}

type ExtendedEntities struct {
	Media []Media `firestore:"media" json:"media"`
}

type Hashtag struct {
	Indices []int  `firestore:"indices" json:"indices"`
	Text    string `firestore:"text"    json:"text"`
}

type URL struct {
	DisplayURL  string   `firestore:"displayUrl"  json:"display_url"`
	ExpandedURL string   `firestore:"expandedUrl" json:"expanded_url"`
	Indices     []int    `firestore:"indices"     json:"indices"`
	URL         string   `firestore:"url"         json:"url"`
	Unwound     *Unwound `firestore:"unwound"     json:"unwound"`
}

type Unwound struct {
	URL         string `firestore:"url"         json:"url"`
	Status      int    `firestore:"status"      json:"status"`
	Title       string `firestore:"title"       json:"title"`
	Description string `firestore:"description" json:"description"`
}

type UserMention struct {
	// [Unused] ID         int64  `firestore:"id"         json:"id"`
	IDStr      string `firestore:"idStr"      json:"id_str"`
	Indices    []int  `firestore:"indices"    json:"indices"`
	Name       string `firestore:"name"       json:"name"`
	ScreenName string `firestore:"screenName" json:"screen_name"`
}

type Symbol struct {
	Indices []int  `firestore:"indices" json:"indices"`
	Text    string `firestore:"text"    json:"text"`
}

type Poll struct {
	Options     []Option  `firestore:"options"     json:"options"`
	EndDatetime time.Time `firestore:"endDatetime" json:"end_datetime"`
	// [Unused] DurationMinutes int
}

type Option struct {
	Position int    `firestore:"position" json:"position"`
	Text     string `firestore:"text"     json:"text"`
}
