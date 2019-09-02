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

func (l *ListenNotes) BestPodcasts(options PodcastsOptions) (podcasts Podcasts, err error) {
	return bestPodcastsRequest(l.httpClient, l.ApiKey, options)
}
