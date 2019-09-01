package listennotes

import "net/http"

type ListenNotes struct {
	ApiKey     string
	URL        string
	httpClient *http.Client
}

var (
	listenNotesBaseURL = "https://listen-api.listennotes.com/api/v2/"
)

func NewListenNotesClient(key string) *ListenNotes {
	return &ListenNotes{
		ApiKey:     key,
		URL:        listenNotesBaseURL,
		httpClient: &http.Client{},
	}
}

func (l *ListenNotes) BestPodcasts(genre_id string, region string, page int, safe_mode int) {
	bestPodcastsRequest(l.httpClient, listenNotesBaseURL, genre_id, region, page, safe_mode)
}
