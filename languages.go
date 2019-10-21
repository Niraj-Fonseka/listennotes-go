package listennotes

import (
	"encoding/json"
	"net/http"
)

type Languages struct {
	Languages []string `json:"languages"`
}

var (
	languagesURL = "languages"
)

func getLanguages(client *http.Client, token string) (languages []string, err error) {

	podcastsURL := listenNotesBaseURL + languagesURL

	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return languages, err
	}

	var languagesResp Languages
	err = json.Unmarshal(response, &languagesResp)

	if err != nil {
		return languages, err
	}

	return languagesResp.Languages, nil

}
