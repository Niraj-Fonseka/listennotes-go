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

// BestPodcasts Get a list of curated best podcasts by genre.
// You can get the genre ids from `GET /genres` endpoint.
func (l *ListenNotes) BestPodcasts() (podcasts Podcasts, err error) {
	podcasts, err = bestPodcastsRequest(l.httpClient, l.APIkey, l.Params)
	l.Params = nil

	return podcasts, err
}

// PodcastsByID Fetch detailed meta data and episodes for a specific podcast (up to 10 episodes each time).
// You can use the **next_episode_pub_date** parameter to do pagination and fetch more episodes.
func (l *ListenNotes) PodcastsByID(id string) (podcast Podcast, err error) {
	podcast, err = getPodcastByIDRequest(l.httpClient, l.APIkey, id, l.Params)
	l.Params = nil

	return podcast, err
}

// Genres returns a list of genre
func (l *ListenNotes) Genres() (genre Genres, err error) {
	return getGenreRequest(l.httpClient, l.APIkey)
}

// DetailedMetadataForEpisode fetches detailed meta data for a specific episode.
func (l *ListenNotes) DetailedMetadataForEpisode(id string) (episode EpisodeMetadataDetailed, err error) {
	return getDetailedEpisodeMetaDataByIDRequest(l.httpClient, l.APIkey, id)
}

// BasicEpsiodeMetadata batch fetches basic meta data for up to 10 episodes.
func (l *ListenNotes) BasicEpsiodeMetadata(ids ...string) (episode EpisodeMetadataBasic, err error) {
	episode, err = getBasicEpisodeMetaDataByIDRequest(l.httpClient, l.APIkey, ids)
	l.Params = nil //make this a defer function

	return episode, err
}

//BasicPodcastMetadata batch fetches basic meta data for up to 10 podcasts
func (l *ListenNotes) BasicPodcastMetadata(ids ...string) (podcast PodcastMedataBasic, err error) {
	podcast, err = getBasicPodcastMetaDataByIDRequest(l.httpClient, l.APIkey, ids, l.Params)
	l.Params = nil
	return podcast, err
}

// CuratedPodcastsByID gets detailed meta data of all podcasts in a specific curated list.
func (l *ListenNotes) CuratedPodcastsByID(id string) (podcasts CuratedPodcasts, err error) {
	podcasts, err = getCuratedPodcastListByID(l.httpClient, l.APIkey, id)
	l.Params = nil
	return podcasts, err
}

// CuratedPodcasts A bunch of curated lists from online media. For each list, you'll get basic info of up to 5 podcasts.
func (l *ListenNotes) CuratedPodcasts() (podcasts CuratedPodcastsPage, err error) {
	return getCuratedPodcastList(l.httpClient, l.APIkey, l.Params)
}

// Regions returns a map of country codes (e.g., us, gb...) & country names (United States, United Kingdom...).
func (l *ListenNotes) Regions() (reigons map[string]string, err error) {
	return getRegions(l.httpClient, l.APIkey)
}

// Languages fetches a list of languages that are supported in Listen Notes database.
func (l *ListenNotes) Languages() (languages []string, err error) {
	return getLanguages(l.httpClient, l.APIkey)
}

// JustListen fetch a random podcast episode
func (l *ListenNotes) JustListen() (podcast Episode, err error) {
	return getJustListenPodcastEpisode(l.httpClient, l.APIkey)
}

// PodcastRecommendations fetch up to 8 podcast recommendations based on the given podcast id.
func (l *ListenNotes) PodcastRecommendations(id string) (recommendations PodcastRecommendations, err error) {
	return getPodcastRecommendations(l.httpClient, l.APIkey, id, l.Params)
}

// EpisodeRecommendations fetch up to 8 episode recommendations based on the given episode id.
func (l *ListenNotes) EpisodeRecommendations(id string) (recommendations EpisodeRecommendations, err error) {
	return getEpisodeRecommendations(l.httpClient, l.APIkey, id, l.Params)
}
