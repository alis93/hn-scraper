package hackernews

import (
	"encoding/json"
	"fmt"
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
	ID       int
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

// Calculates and returns how many points this RawItem has
func (s RawItem) calculatePoints() int {
	return s.Score
}

// Calculates and returns how many comments this RawItem has
func (s RawItem) countComments() int {
	return s.Descendants
}

// NOTE: used value receiver as we are not mutating
// Converts a RawItem into a Story struct
// Only works if the ItemType is story.
func (item RawItem) ConvertToStory(rank int, opts *StoryOpts) (*Story, error) {
	expectedType := "story"
	if item.ItemType != expectedType {
		return nil, &InvalidItemTypeErr{expectedType, item.ItemType}
	}
	if rank <= 0 {
		return nil, &MinValErr{1, rank}
	}
	validTitle, err := opts.ValidateStr(item.Title)
	if err != nil {
		return nil, err
	}
	validAuthor, err := opts.ValidateStr(item.By)
	if err != nil {
		return nil, err
	}
	if !isValidURLScheme(item.URL) {
		return nil, fmt.Errorf("item has an invalid url scheme. \t %s", item.URL)
	}
	points := item.calculatePoints()
	if opts.minPoints > points {
		return nil, &MinValErr{opts.minPoints, points}
	}
	comments := item.countComments()
	if opts.minComments > comments {
		return nil, &MinValErr{opts.minComments, comments}
	}
	story := &Story{
		ID:       item.GetID(),
		Title:    validTitle,
		URL:      item.URL,
		Author:   validAuthor,
		Points:   points,
		Comments: comments,
		Rank:     rank,
	}
	return story, nil
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
