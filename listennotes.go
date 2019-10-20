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
	listenNotesBaseURL = "https://listen-api.listennotes.com/api/v2/"
)

// NewListenNotesClient creates a new ListenNotes instance
func NewListenNotesClient(key string) *ListenNotes {

	params := make(Params)
	return &ListenNotes{
		APIkey:     key,
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

// PodcastsByID returns a podcast given ID
func (l *ListenNotes) PodcastsByID(id string) (podcast Podcast, err error) {
	podcast, err = getPodcastByIDRequest(l.httpClient, l.APIkey, id, l.Params)
	l.Params = nil

	return podcast, err
}

// Genres returns a list of genre
func (l *ListenNotes) Genres() (genre Genres, err error) {
	return getGenreRequest(l.httpClient, l.APIkey)
}

func (l *ListenNotes) GetEpisodeMetadataByID(id string) (episode EpisodeMetadata, err error) {
	return getEpisodeMetaDataByIDRequest(l.httpClient, l.APIkey, id)
}
