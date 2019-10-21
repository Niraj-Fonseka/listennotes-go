package listennotes

import (
	"encoding/json"
	"net/http"
)

type RandomPodcast struct {
	ID                 string `json:"id"`
	Audio              string `json:"audio"`
	Image              string `json:"image"`
	Title              string `json:"title"`
	Publisher          string `json:"publisher"`
	Thumbnail          string `json:"thumbnail"`
	PodcastID          string `json:"podcast_id"`
	Description        string `json:"description"`
	PubDateMs          int64  `json:"pub_date_ms"`
	PodcastTitle       string `json:"podcast_title"`
	ListennotesURL     string `json:"listennotes_url"`
	AudioLengthSec     int    `json:"audio_length_sec"`
	ExplicitContent    bool   `json:"explicit_content"`
	MaybeAudioInvalid  bool   `json:"maybe_audio_invalid"`
	ListennotesEditURL string `json:"listennotes_edit_url"`
}

var (
	justListenURL = "just_listen"
)

func getJustListenPodcast(client *http.Client, token string) (podcast RandomPodcast, err error) {

	podcastsURL := listenNotesBaseURL + justListenURL

	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return podcast, err
	}

	var podcastResp RandomPodcast

	err = json.Unmarshal(response, &podcastResp)

	if err != nil {
		return podcast, err
	}

	return podcastResp, nil

}
