package listennotes

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// Params is the type to store all the query params
type Params map[string]string

func newGetRequest(url string, token string, client *http.Client) ([]byte, error) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-type", "application/json")
	request.Header.Set("X-ListenAPI-Key", token)

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

func paramValidator(params Params, needParams bool) error {
	if needParams && len(params) <= 0 {
		return errors.New("No parameters provided")
	}

	return nil
}

func buildURL(params Params) string {

	buf := bytes.Buffer{}
	for k, v := range params {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		buf.WriteString("&")
	}

	return buf.String()
}
