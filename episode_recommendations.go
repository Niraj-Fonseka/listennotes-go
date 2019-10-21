package listennotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EpisodeRecommendations struct {
	Recommendations []EpisodeRecomendation `json:"recommendations"`
}

type EpisodeRecomendation struct {
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
	episodesRecommendationsURL = "episodes/%s/recommendations"
)

func getEpisodeRecommendations(client *http.Client, token string, id string, options Params) (recommendations EpisodeRecommendations, err error) {

	builtURL := buildURL(options)

	fillID := fmt.Sprintf(episodesRecommendationsURL, id)
	podcastsURL := listenNotesBaseURL + fillID + builtURL

	fmt.Println(podcastsURL)
	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return recommendations, err
	}

	var recommendationResp EpisodeRecommendations

	err = json.Unmarshal(response, &recommendationResp)

	if err != nil {
		return recommendations, err
	}

	return recommendationResp, nil

}
