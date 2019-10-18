package podcasts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getPodcastByIDRequest(client *http.Client, token string, id string, options PodcastsOptions) (podcast Podcast, err error) {
	podcastByID := fmt.Sprintf(pocastByIDURL, listenNotesBaseURL, id, options.NextEpisodePublishDate, options.Sort)

	fmt.Println(podcastByID)
	response, err := newGetRequest(podcastByID, token, client)

	if err != nil {
		return podcast, err
	}

	var podcastResp Podcast

	err = json.Unmarshal(response, &podcastResp)

	if err != nil {
		return podcast, err
	}

	return podcastResp, nil
}
