package listennotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EpisodeRecommendations struct {
	Recommendations []Episode `json:"recommendations"`
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
