package listennotes

import (
	"encoding/json"
	"net/http"
)

// Genres stores a list of genre
type Genres struct {
	Genres []Genre `json:"genres"`
}

// Genre stores a genre
type Genre struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
}

var (
	genresURL = "genres"
)

func getGenreRequest(client *http.Client, token string) (genres Genres, err error) {

	podcastsURL := listenNotesBaseURL + genresURL

	response, err := newGetRequest(podcastsURL, token, client)

	if err != nil {
		return genres, err
	}

	var genresResp Genres

	err = json.Unmarshal(response, &genresResp)

	if err != nil {
		return genres, err
	}

	return genresResp, nil

}
