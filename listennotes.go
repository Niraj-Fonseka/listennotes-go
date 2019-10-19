package listennotes

import "net/http"

// ListenNotes holds the configurations for the listen notes client
type ListenNotes struct {
	APIkey     string
	URL        string
	Params     Params
	httpClient *http.Client
}

var (
	// ListenNotesBaseURL holds the base url for the listen notes api
	ListenNotesBaseURL = "https://listen-api.listennotes.com/api/v2/"
)

// NewListenNotesClient creates a new ListenNotes instance
func NewListenNotesClient(key string) *ListenNotes {

	params := make(Params)
	return &ListenNotes{
		APIkey:     key,
		URL:        listenNotesBaseURL,
		httpClient: &http.Client{},
		Params:     params,
	}
}

// BestPodcasts returns the best podcasts at the moment
func (l *ListenNotes) BestPodcasts() (podcasts Podcasts, err error) {
	podcasts, err = bestPodcastsRequest(l.httpClient, l.APIkey, l.Params)
	l.Params = nil

	return podcasts, err
}

// func (l *ListenNotes) PodcastsByID(id string, options PodcastsOptions) (podcast Podcast, err error) {
// 	return getPodcastByIDRequest(l.httpClient, l.ApiKey, id, options)
// }
