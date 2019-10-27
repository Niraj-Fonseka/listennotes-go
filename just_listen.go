package listennotes

import (
	"encoding/json"
	"net/http"
)

var (
	justListenURL = "just_listen"
)

func getJustListenPodcastEpisode(client *http.Client, token string) (epsiode Episode, err error) {

	podcastsURL := listenNotesBaseURL + justListenURL

	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return epsiode, err
	}

	var podcastResp Episode

	err = json.Unmarshal(response, &podcastResp)

	if err != nil {
		return epsiode, err
	}

	return podcastResp, nil

}
