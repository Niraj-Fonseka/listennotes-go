package listennotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EpisodeMetadataDetailed struct {
	ID      string `json:"id"`
	Audio   string `json:"audio"`
	Image   string `json:"image"`
	Title   string `json:"title"`
	Podcast struct {
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
	} `json:"podcast"`
	Thumbnail          string `json:"thumbnail"`
	Description        string `json:"description"`
	PubDateMs          int64  `json:"pub_date_ms"`
	ListennotesURL     string `json:"listennotes_url"`
	AudioLengthSec     int    `json:"audio_length_sec"`
	ExplicitContent    bool   `json:"explicit_content"`
	MaybeAudioInvalid  bool   `json:"maybe_audio_invalid"`
	ListennotesEditURL string `json:"listennotes_edit_url"`
}

func getDetailedEpisodeMetaDataByIDRequest(client *http.Client, token string, id string) (episode EpisodeMetadataDetailed, err error) {

	epsiodesMetadataURL := listenNotesBaseURL + episodesURL + id

	fmt.Println(epsiodesMetadataURL)
	response, err := newGetRequest(epsiodesMetadataURL, token, client)

	if err != nil {
		return episode, err
	}

	var episodeMetadata EpisodeMetadataDetailed

	err = json.Unmarshal(response, &episodeMetadata)

	if err != nil {
		return episode, err
	}

	return episodeMetadata, nil
}

// func csvIDs(ids []string) string {
// 	buf := bytes.Buffer{}
// 	buf.WriteString(ids[0])

// 	if len(ids) > 1 {
// 		for i := range ids[1:] {
// 			buf.WriteString(",")
// 			buf.WriteString(ids[i])
// 		}
// 	}
// 	return buf.String()
// }
