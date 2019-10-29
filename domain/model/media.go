package model

type Media struct {
	DisplayURL  string `firestore:"displayUrl"        json:"display_url"`
	ExpandedURL string `firestore:"expandedUrl"       json:"expanded_url"`
	// [Unused] ID                int64   `firestore:"id"                json:"id"`
	IDStr         string `firestore:"idStr"             json:"id_str"`
	Indices       []int  `firestore:"indices"           json:"indices"`
	MediaURL      string `firestore:"mediaUrl"          json:"media_url"`
	MediaURLHTTPS string `firestore:"mediaUrlHttps"     json:"media_url_https"`
	Sizes         Sizes  `firestore:"sizes"             json:"sizes"`
	// [Unused] SourceStatusID    *int64  `firestore:"sourceStatusId"    json:"source_status_id"`
	SourceStatusIDStr *string `firestore:"sourceStatusIdStr" json:"source_status_id_str"`
	Type              string  `firestore:"type"              json:"type"`
	URL               string  `firestore:"url"               json:"url"`
}

type Sizes struct {
	Thumb  Size `firestore:"thumb"  json:"thumb"`
	Large  Size `firestore:"large"  json:"large"`
	Medium Size `firestore:"medium" json:"medium"`
	Small  Size `firestore:"small"  json:"small"`
}

type Size struct {
	W      int    `firestore:"w"      json:"w"`
	H      int    `firestore:"h"      json:"h"`
	Resize string `firestore:"resize" json:"resize"`
}
