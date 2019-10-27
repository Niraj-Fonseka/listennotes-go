package listennotes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type EpisodeMetadataBasic struct {
	Episodes []Episode `json:"episodes"`
}

type Episode struct {
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
	episodesURL = "episodes/"
)

func getBasicEpisodeMetaDataByIDRequest(client *http.Client, token string, ids []string) (episode EpisodeMetadataBasic, err error) {

	epsiodesMetadataURL := listenNotesBaseURL + episodesURL

	//build csv string

	if len(ids) <= 0 {
		return episode, errors.New("No episode id's provided")
	}

	csvIDs := csvIDs(ids)

	fmt.Println("CSVIds ", csvIDs)

	p := Params{}

	p["ids"] = csvIDs
	response, err := newPostRequest(epsiodesMetadataURL, token, client, buildFormValues(p))

	if err != nil {
		return episode, err
	}

	var episodeMetadata EpisodeMetadataBasic

	err = json.Unmarshal(response, &episodeMetadata)

	if err != nil {
		return episode, err
	}

	return episodeMetadata, nil

}

func csvIDs(ids []string) string {
	buf := bytes.Buffer{}
	buf.WriteString(ids[0])
	fmt.Println("added ", ids[0])
	if len(ids) > 1 {
		for _, v := range ids[1:] {
			buf.WriteString(",")
			buf.WriteString(v)

		}
	}
	return buf.String()
}

func buildFormValues(params Params) url.Values {
	data := url.Values{}
	for k, v := range params {
		data.Add(k, v)
	}
	return data
}
