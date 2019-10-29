package model

import (
	"encoding/json"
	"time"

	"github.com/tidwall/gjson"
)

type User struct {
	// [Unused] ID          int64         `firestore:"id"          json:"id"`
	IDStr       string        `firestore:"idStr"       json:"id_str"`
	Name        string        `firestore:"name"        json:"name"`
	ScreenName  string        `firestore:"screenName"  json:"screen_name"`
	Location    *string       `firestore:"location"    json:"location"`
	URL         *string       `firestore:"url"         json:"url"`
	Entities    *UserEntities `firestore:"entities"    json:"entities"`
	Description *string       `firestore:"description" json:"description"`
	// [Unused] Drived      []Enrichment
	Protected       bool      `firestore:"protected"       json:"protected"`
	Verified        bool      `firestore:"verified"        json:"verified"`
	FollowersCount  int       `firestore:"followersCount"  json:"followers_count"`
	FriendsCount    int       `firestore:"friendsCount"    json:"friends_count"`
	ListedCount     int       `firestore:"listedCount"     json:"listed_count"`
	FavouritesCount int       `firestore:"favouritesCount" json:"favourites_count"`
	StatusesCount   int       `firestore:"statusesCount"   json:"statuses_count"`
	CreatedAt       time.Time `firestore:"createdAt"       json:"-"`
	// [Unused] UTCOffset       nil
	// [Unused] TimeZone        nil
	// [Unused] GEOEnabled                     bool
	Lang string `firestore:"lang" json:"lang"`
	// [Unused] ContributorsEnabled            bool
	// [Unused] ProfileBackgroundColor         string
	// [Unused] ProfileBackgroundImageURL      string
	// [Unused] ProfileBackgroundImageURLHTTPS string
	// [Unused] ProfileBackgroundTitle         string
	// [Unused] ProfileBannerURL               string
	// [Unused] ProfileImageURL                string
	ProfileImageURLHTTPS string `firestore:"profileImageUrlHttps" json:"profile_image_url_https"`
	// [Unused] ProfileLinkColor               string
	// [Unused] ProfileSidebarBorderColor      string
	// [Unused] ProfileSidebarFillColor        string
	// [Unused] ProfileTextColor               string
	// [Unused] ProfileUseBackgroundImage      bool
	// [Unused] DefaultProfile                 bool
	// [Unused] DefaultProfileImage            bool
	// [Unused] WithheldInContries             []string
	// [Unused] WithheldScope                  string
}

func (u *User) UnmarshalJSON(data []byte) error {
	type u2 User
	if err := json.Unmarshal(data, (*u2)(u)); err != nil {
		return err
	}
	jsonS := string(data)
	createdAt := gjson.Get(jsonS, "created_at").String()
	tm, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", createdAt)
	if err != nil {
		return err
	}
	u.CreatedAt = tm
	return nil
}

type UserEntities struct {
	URL struct {
		URLS []URL `firestore:"urls" json:"urls"`
	} `firestore:"url" json:"url"`
}
