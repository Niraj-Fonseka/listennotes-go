package listennotes

import (
	"encoding/json"
	"net/http"
)

type CuratedPodcasts struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Podcasts []struct {
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
	} `json:"podcasts"`
	SourceURL      string `json:"source_url"`
	Description    string `json:"description"`
	PubDateMs      int64  `json:"pub_date_ms"`
	SourceDomain   string `json:"source_domain"`
	ListennotesURL string `json:"listennotes_url"`
}

var (
	curatedPodcastsIDURL = "curated_podcasts/"
)

func getCuratedPodcastListByID(client *http.Client, token string, id string) (podcasts CuratedPodcasts, err error) {

	podcastsURL := listenNotesBaseURL + curatedPodcastsIDURL + id

	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return podcasts, err
	}

	var curatedResponse CuratedPodcasts

	err = json.Unmarshal(response, &curatedResponse)

	if err != nil {
		return podcasts, err
	}

	return curatedResponse, nil

}
