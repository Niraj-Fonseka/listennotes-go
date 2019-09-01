package listennotes

import (
	"io/ioutil"
	"net/http"
)

func newGetRequest(url string, client *http.Client) ([]byte, error) {

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
