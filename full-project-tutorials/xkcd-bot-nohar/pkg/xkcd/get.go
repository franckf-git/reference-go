package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// Latest is passed to the Get function to get the latest comic
	Latest = "latest"

	urlTemplate = "https://xkcd.com/%s/info.0.json"
	urlLatest   = "https://xkcd.com/info.0.json"
)

// Get retrieves the metadata from an xkcd given its unique number.
// If num is xkcd.Latest, the metadata from the latest comic is returned.
func Get(num string) (*Comic, error) {
	var url string
	if num == Latest {
		url = urlLatest
	} else {
		url = fmt.Sprintf(urlTemplate, num)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("couldn't get '%s': %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("couldn't get '%s': %s", url, resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read body at '%s': %w", url, err)
	}

	var c Comic
	if err = json.Unmarshal(body, &c); err != nil {
		return nil, fmt.Errorf("couldn't parse data from '%s': %w", url, err)
	}
	return &c, nil
}
