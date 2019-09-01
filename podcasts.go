package listennotes

import (
	"fmt"
	"net/http"
)

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
	podcastsURL = "%s/best_podcasts?genre_id=%d&page=%d&region=%s&safe_mode=%d"
)

func bestPodcastsRequest(client *http.Client, base_url string, genre_id string, region string, page int, safe_mode int) {
	podcastsURL := fmt.Sprintf(podcastsURL, base_url, genre_id, page, region, safe_mode)

	response, err := newGetRequest(podcastsURL, client)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(response))

}
