package listennotes

import (
	"net/http"
)

type ListenNotes struct {
	ApiKey     string
	URL        string
	httpclient *http.Client
}

var (
	listenNotesBaseURL = "https://listen-api.listennotes.com/api/v2/"
)

func NewListenNotesClient(key string) *ListenNotes {

	return &ListenNotes{
		ApiKey:     key,
		URL:        listenNotesBaseURL,
		httpclient: &http.Client{},
	}
}
