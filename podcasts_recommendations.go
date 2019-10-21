package listennotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PodcastRecommendations struct {
	Recommendations []PodcastRecommendation `json:"recommendations"`
}

type PodcastRecommendation struct {
	ID    string `json:"id"`
	Rss   string `json:"rss"`
	Email string `json:"email"`
	Extra struct {
		URL1            string `json:"url1"`
		URL2            string `json:"url2"`
		URL3            string `json:"url3"`
		GoogleURL       string `json:"google_url"`
		SpotifyURL      string `json:"spotify_url"`
		YoutubeURL      string `json:"youtube_url"`
		LinkedinURL     string `json:"linkedin_url"`
		WechatHandle    string `json:"wechat_handle"`
		PatreonHandle   string `json:"patreon_handle"`
		TwitterHandle   string `json:"twitter_handle"`
		FacebookHandle  string `json:"facebook_handle"`
		InstagramHandle string `json:"instagram_handle"`
	} `json:"extra"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Country     string `json:"country"`
	Website     string `json:"website"`
	Language    string `json:"language"`
	GenreIds    []int  `json:"genre_ids"`
	ItunesID    int    `json:"itunes_id"`
	Publisher   string `json:"publisher"`
	Thumbnail   string `json:"thumbnail"`
	IsClaimed   bool   `json:"is_claimed"`
	Description string `json:"description"`
	LookingFor  struct {
		Guests         bool `json:"guests"`
		Cohosts        bool `json:"cohosts"`
		Sponsors       bool `json:"sponsors"`
		CrossPromotion bool `json:"cross_promotion"`
	} `json:"looking_for"`
	TotalEpisodes     int    `json:"total_episodes"`
	ListennotesURL    string `json:"listennotes_url"`
	ExplicitContent   bool   `json:"explicit_content"`
	LatestPubDateMs   int64  `json:"latest_pub_date_ms"`
	EarliestPubDateMs int64  `json:"earliest_pub_date_ms"`
}

var (
	podcastsRecommendationsURL = "podcasts/%s/recommendations"
)

func getPodcastRecommendations(client *http.Client, token string, id string, options Params) (recommendations PodcastRecommendations, err error) {

	builtURL := buildURL(options)

	fillID := fmt.Sprintf(podcastsRecommendationsURL, id)
	podcastsURL := listenNotesBaseURL + fillID + builtURL

	fmt.Println(podcastsURL)
	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return recommendations, err
	}

	var recommendationResp PodcastRecommendations

	err = json.Unmarshal(response, &recommendationResp)

	if err != nil {
		return recommendations, err
	}

	return recommendationResp, nil

}
