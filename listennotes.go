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

func (l *ListenNotes) DetailedMetadataForEpisode(id string) (episode EpisodeMetadataDetailed, err error) {
	return getDetailedEpisodeMetaDataByIDRequest(l.httpClient, l.APIkey, id)
}

func (l *ListenNotes) BasicMetadataForEpisodes(ids ...string) (episode EpisodeMetadataBasic, err error) {
	episode, err = getBasicEpisodeMetaDataByIDRequest(l.httpClient, l.APIkey, ids)
	l.Params = nil //make this a defer function

	return episode, err
}

func (l *ListenNotes) BasicPodcastMetaDataByIDRequest(ids ...string) (podcast PodcastMedataBasic, err error) {
	podcast, err = getBasicPodcastMetaDataByIDRequest(l.httpClient, l.APIkey, ids, l.Params)
	l.Params = nil
	return podcast, err
}

func (l *ListenNotes) CuratedPodcastsList(id string) (podcasts CuratedPodcasts, err error) {
	podcasts, err = getCuratedPodcastList(l.httpClient, l.APIkey, id)
	l.Params = nil
	return podcasts, err
}

func (l *ListenNotes) Regions() (reigons map[string]string, err error) {
	return getRegions(l.httpClient, l.APIkey)
}

func (l *ListenNotes) Languages() (languages []string, err error) {
	return getLanguages(l.httpClient, l.APIkey)
}

func (l *ListenNotes) JustListen() (podcast RandomPodcast, err error) {
	return getJustListenPodcast(l.httpClient, l.APIkey)
}
