package hackernews

import (
	"encoding/json"
)

// Note: This interface may not be needed currently
type Item interface {
	GetID() int
}

// Represents a hackernews item retrieved from the api.
type RawItem struct {
	ID          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	ItemType    string `json:"type"` //we should use enum for type of item
	By          string `json:"by"`
	Timestamp   int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Poll        int    `json:"poll"`
	Kids        []int  `json:"kids"`
	URL         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Parts       []int  `json:"parts"`
	Descendants int    `json:"descendants"`
}

// Represents a Story item
type Story struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	URL      string `json:"uri"`
	Author   string `json:"author"`
	Points   int    `json:"points"`
	Comments int    `json:"comments"`
	Rank     int    `json:"rank"`
}

// Returns the id of the item
func (item RawItem) GetID() int {
	return item.ID
}

// Returns the id of the story.
func (s Story) getID() int {
	return s.ID
}

// Converts the story into a string representation
func (s Story) String() string {
	prettyJson, _ := json.MarshalIndent(s, "", "    ")
	return string(prettyJson)
}
