package listennotes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type PodcastMetaData struct {
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
type PodcastMedataBasic struct {
	Podcasts       []PodcastMetaData      `json:"podcasts"`
	LatestEpisodes []BasicEpisodeMetadata `json:"latest_episodes"`
}

func getBasicPodcastMetaDataByIDRequest(client *http.Client, token string, ids []string, params Params) (podcast PodcastMedataBasic, err error) {

	epsiodesMetadataURL := listenNotesBaseURL + podcastURL

	//build csv string

	if len(ids) <= 0 {
		return podcast, errors.New("No podcast id's provided")
	}

	csvIDs := csvIDs(ids)

	fmt.Println("CSVIds ", csvIDs)

	params["ids"] = csvIDs

	response, err := newPostRequest(epsiodesMetadataURL, token, client, buildFormValues(params))

	if err != nil {
		return podcast, err
	}

	var podcastMetadata PodcastMedataBasic

	err = json.Unmarshal(response, &podcastMetadata)

	if err != nil {
		return podcast, err
	}

	return podcastMetadata, nil

}
