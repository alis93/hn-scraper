package hackernews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	BASE_URL             = "https://hacker-news.firebaseio.com"
	API_VERSION          = "v0"
	TOP_STORIES_ENDPOINT = "topstories.json"
	ITEM_ENDPOINT        = "item/%d.json"
)

type Client struct {
	http   *http.Client
	apiURL string
}

// Creates a client with given timeout
func NewClient(timeout int) (*Client, error) {

	apiURL := fmt.Sprintf("%s/%s", BASE_URL, API_VERSION)

	if !isValidURLScheme(apiURL) {
		return nil, &ClientErr{"API URL is an invalid URL"}
	}

	if timeout <= 0 {
		return nil, InvalidTimeOutErr
	}

	return &Client{
		&http.Client{Timeout: time.Duration(timeout) * time.Second},
		apiURL,
	}, nil
}

// Returns list of top n stories on hackernews, where n is amount
// Amount must be between 1 and 500 inclusive
// Returns error if it fails
func (c Client) GetTopStoryIds(amount int) ([]int, error) {

	if amount > 500 || amount <= 0 {
		return nil, OutOfRangeErr
	}

	// api endpoint
	endpoint := fmt.Sprintf("%s/%s", c.apiURL, TOP_STORIES_ENDPOINT)

	// send get request
	res, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	// read response as []byte
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var storyList []int

	// convert from []byte into []int storing into storylist
	if err := json.Unmarshal(body, &storyList); err != nil {
		return nil, err
	}

	if amount > len(storyList) {
		return storyList, nil
	}

	return storyList[:amount], nil

}

// Retrieves the item from hackerrank using the id passed in.
func (c Client) GetItem(id int) (*RawItem, error) {
	endpoint := fmt.Sprintf("%s/%s", c.apiURL, ITEM_ENDPOINT)
	endpoint = fmt.Sprintf(endpoint, id)
	res, err := c.http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	// read response into []byte
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	item := &RawItem{}

	// convert body into RawItem struct. Uses the Json tags defined on struct.
	if err := json.Unmarshal(body, item); err != nil {
		return nil, err
	}

	return item, nil

}
