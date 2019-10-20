package listennotes

// func getEpisodesByIDRequest(client *http.Client, token string, id string, options Params) (podcast Podcast, err error) {
// 	podcastByID := fmt.Sprintf(pocastByIDURL, ListenNotesBaseURL, id, options["next_episode_pub_date"], options["sort"])

// 	response, err := newGetRequest(podcastByID, token, client)

// 	if err != nil {
// 		return podcast, err
// 	}

// 	var podcastResp Podcast

// 	err = json.Unmarshal(response, &podcastResp)

// 	if err != nil {
// 		return podcast, err
// 	}

// 	return podcastResp, nil
// }
