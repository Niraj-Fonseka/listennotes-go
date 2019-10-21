package listennotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CuratedPodcastsPage struct {
	Total        int  `json:"total"`
	HasNext      bool `json:"has_next"`
	PageNumber   int  `json:"page_number"`
	HasPrevious  bool `json:"has_previous"`
	CuratedLists []struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Podcasts []struct {
			ID             string `json:"id"`
			Image          string `json:"image"`
			Title          string `json:"title"`
			Publisher      string `json:"publisher"`
			Thumbnail      string `json:"thumbnail"`
			ListennotesURL string `json:"listennotes_url"`
		} `json:"podcasts"`
		SourceURL      string `json:"source_url"`
		Description    string `json:"description"`
		PubDateMs      int64  `json:"pub_date_ms"`
		SourceDomain   string `json:"source_domain"`
		ListennotesURL string `json:"listennotes_url"`
	} `json:"curated_lists"`
	NextPageNumber     int `json:"next_page_number"`
	PreviousPageNumber int `json:"previous_page_number"`
}

var (
	curatedPodcastsURL = "curated_podcasts?"
)

func getCuratedPodcastList(client *http.Client, token string, options Params) (podcasts CuratedPodcastsPage, err error) {

	builtURL := buildURL(options)

	podcastsURL := listenNotesBaseURL + curatedPodcastsURL + builtURL

	fmt.Println(podcastsURL)
	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return podcasts, err
	}

	var curatedResponse CuratedPodcastsPage

	err = json.Unmarshal(response, &curatedResponse)

	if err != nil {
		return podcasts, err
	}

	return curatedResponse, nil

}
