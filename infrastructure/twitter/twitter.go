package twitter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/dappface/news-loader/domain/model"
	"github.com/dghubble/oauth1"
)

var Client *client

func init() {
	at := os.Getenv("TWITTER_ACCESS_TOKEN")
	if at == "" {
		log.Fatalln("Twitter access token is missing")
	}

	ts := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	if ts == "" {
		log.Fatalln("Twitter access token secret is missing")
	}

	ak := os.Getenv("TWITTER_API_KEY")
	if ak == "" {
		log.Fatalln("Twitter api key is missing")
	}

	as := os.Getenv("TWITTER_API_SECRET")
	if as == "" {
		log.Fatalln("Twitter api secret is missing")
	}

	config := oauth1.NewConfig(ak, as)
	token := oauth1.NewToken(at, ts)

	c := config.Client(oauth1.NoContext, token)
	Client = newClient(c, "https://api.twitter.com/1.1")
}

func newClient(c *http.Client, basePath string) *client {
	return &client{
		c,
		basePath,
	}
}

type client struct {
	c        *http.Client
	basePath string
}

// func (s *client) GetTweet(ID int64) (*model.Tweet, error) {
// 	values := url.Values{}
// 	values.Add("id", strconv.FormatInt(ID, 10))
// 	path := s.basePath + "/statuses/show.json" + "?" + values.Encode()
// 	resp, err := s.c.Get(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var t model.Tweet
// 	if err = json.Unmarshal(body, &t); err != nil {
// 		return nil, err
// 	}

// 	return &t, nil
// }

func (s *client) GetTweetsBySlug(slug string, ownerScreenName string, sinceID *string) (*[]model.Tweet, error) {
	values := url.Values{}
	values.Add("slug", slug)
	values.Add("owner_screen_name", ownerScreenName)
	if sinceID != nil {
		values.Add("since_id", *sinceID)
	}
	path := s.basePath + "/lists/statuses.json" + "?" + values.Encode()
	resp, err := s.c.Get(path)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ts []model.Tweet
	if err = json.Unmarshal(body, &ts); err != nil {
		return nil, err
	}

	return &ts, nil
}
