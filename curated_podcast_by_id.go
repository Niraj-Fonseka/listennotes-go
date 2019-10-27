package listennotes

import (
	"encoding/json"
	"net/http"
)

type CuratedPodcasts struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Podcasts       []Podcast `json:"podcasts"`
	SourceURL      string    `json:"source_url"`
	Description    string    `json:"description"`
	PubDateMs      int64     `json:"pub_date_ms"`
	SourceDomain   string    `json:"source_domain"`
	ListennotesURL string    `json:"listennotes_url"`
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
