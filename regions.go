package listennotes

import (
	"encoding/json"
	"net/http"
)

type Regions struct {
	Regions map[string]string `json:"regions"`
}

var (
	reigonsURL = "regions"
)

func getRegions(client *http.Client, token string) (regions map[string]string, err error) {

	podcastsURL := listenNotesBaseURL + reigonsURL

	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return regions, err
	}

	var regionsResp Regions
	err = json.Unmarshal(response, &regionsResp)

	if err != nil {
		return regions, err
	}

	return regionsResp.Regions, nil

}
