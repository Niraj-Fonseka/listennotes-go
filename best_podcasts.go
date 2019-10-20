package listennotes

import (
	"encoding/json"
	"net/http"
)

// Podcasts is the response for multiple podcasts
type Podcasts struct {
	PageNumber         int       `json:"page_number"`
	PreviousPageNumber int       `json:"previous_page_number"`
	NextPageNumber     int       `json:"next_page_number"`
	Total              int       `json:"total"`
	Podcasts           []Podcast `json:"podcasts"`
}

// Podcast is the response for one podcast
type Podcast struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	Publisher         string `json:"publisher"`
	Image             string `json:"image"`
	Thumbnail         string `json:"thumbnail"`
	ListennotesURL    string `json:"listennotes_url"`
	TotalEpisodes     int    `json:"total_episodes"`
	ExplicitContent   bool   `json:"explicit_content"`
	Description       string `json:"description"`
	ItunesID          int    `json:"itunes_id"`
	Rss               string `json:"rss"`
	LatestPubDateMs   int64  `json:"latest_pub_date_ms"`
	EarliestPubDateMs int64  `json:"earliest_pub_date_ms"`
	Language          string `json:"language"`
	Country           string `json:"country"`
	Website           string `json:"website"`
	Extra             struct {
		TwitterHandle   string `json:"twitter_handle"`
		FacebookHandle  string `json:"facebook_handle"`
		InstagramHandle string `json:"instagram_handle"`
		WechatHandle    string `json:"wechat_handle"`
		PatreonHandle   string `json:"patreon_handle"`
		YoutubeURL      string `json:"youtube_url"`
		LinkedinURL     string `json:"linkedin_url"`
		SpotifyURL      string `json:"spotify_url"`
		GoogleURL       string `json:"google_url"`
		URL1            string `json:"url1"`
		URL2            string `json:"url2"`
		URL3            string `json:"url3"`
	} `json:"extra"`
	IsClaimed  bool   `json:"is_claimed"`
	Email      string `json:"email"`
	LookingFor struct {
		Sponsors       bool `json:"sponsors"`
		Guests         bool `json:"guests"`
		Cohosts        bool `json:"cohosts"`
		CrossPromotion bool `json:"cross_promotion"`
	} `json:"looking_for"`
	GenreIds []int `json:"genre_ids"`
}

var (
	bestPodcastsURL = "best_podcasts?"
)

func bestPodcastsRequest(client *http.Client, token string, options Params) (podcasts Podcasts, err error) {

	builtURL := buildURL(options)
	podcastsURL := listenNotesBaseURL + bestPodcastsURL + builtURL

	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return podcasts, err
	}

	var podcastResp Podcasts

	err = json.Unmarshal(response, &podcastResp)

	if err != nil {
		return podcasts, err
	}

	return podcastResp, nil

}
