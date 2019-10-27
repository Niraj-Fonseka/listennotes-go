package listennotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PodcastRecommendations struct {
	Recommendations []PodcastRecommendation `json:"recommendations"`
}

type PodcastRecommendation PodcastMetaData

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
