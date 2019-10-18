package listennotes

import "net/http"

type ListenNotes struct {
	ApiKey     string
	URL        string
	httpClient *http.Client
}

var (
	ListenNotesBaseURL = "https://listen-api.listennotes.com/api/v2/"
)

func NewListenNotesClient(key string) *ListenNotes {
	return &ListenNotes{
		ApiKey:     key,
		URL:        listenNotesBaseURL,
		httpClient: &http.Client{},
	}
}

// get best podcasts
func (l *ListenNotes) BestPodcasts(options PodcastsOptions) (podcasts Podcasts, err error) {
	return bestPodcastsRequest(l.httpClient, l.ApiKey, options)
}

func (l *ListenNotes) PodcastsByID(id string, options PodcastsOptions) (podcast Podcast, err error) {
	return getPodcastByIDRequest(l.httpClient, l.ApiKey, id, options)
}
