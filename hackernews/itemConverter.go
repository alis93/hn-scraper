package hackernews

import "fmt"

// type Converter interface {
// 	Convert() Item
// }

// Options for creating a story.
type ItemConverter struct {
	emptyStringsAllowed    bool
	enforceMaxStringLength bool
	maxStringLength        int
	minComments            int
	minPoints              int
}

func NewItemConverter(emptyStringAllowed, enforceMaxStrLength bool, maxStrLength, minComments, minPoints int) (*ItemConverter, error) {
	if enforceMaxStrLength && maxStrLength <= 0 {
		return nil, fmt.Errorf("If enforceMaxStringLength is set, then maxStringLength must be more than 0")
	}
	if minComments <= 0 {
		return nil, &MinValErr{1, minComments}
	}
	if minPoints <= 0 {
		return nil, &MinValErr{1, minPoints}
	}
	return &ItemConverter{
		emptyStringsAllowed:    emptyStringAllowed,
		enforceMaxStringLength: enforceMaxStrLength,
		maxStringLength:        maxStrLength,
		minComments:            minComments,
		minPoints:              minPoints,
	}, nil
}

// Helper function to validate strings.
// Tests and validates string.
// Returns error if invalid.
// Returns string if valid.
// truncates to required length if enforceMaxStringLength and maxStringLength are set
func (cnv ItemConverter) ValidateStr(str string) (string, error) {

	strLen := len(str)
	// test empty strings
	if !cnv.emptyStringsAllowed && strLen <= 0 {
		return "", EmptyStringErr
	}

	finalStr := str

	// test string length
	// if string length more than max, truncate string
	if cnv.enforceMaxStringLength && strLen > cnv.maxStringLength {
		if cnv.maxStringLength <= 0 {
			return "", MaxStringErr
		}
		finalStr = str[:cnv.maxStringLength]
	}

	return finalStr, nil
}

// Calculates and returns how many points this RawItem has.
// if item has less than minpoints then returns 0 and error
func (cnv ItemConverter) calculatePoints(i *RawItem) (int, error) {
	score := i.Score
	if cnv.minPoints > score {
		return 0, &MinValErr{min: cnv.minPoints, actual: score}
	}

	return score, nil
}

// Calculates and returns hoitemitemw many comments this RawItem has
// if item has less than minpoints then returns 0 and error
func (cnv ItemConverter) countComments(i *RawItem) (int, error) {
	numComments := i.Descendants
	if cnv.minComments > numComments {
		return 0, &MinValErr{min: cnv.minComments, actual: numComments}
	}

	return numComments, nil
}

// Converts a RawItem into a Story struct
// Only works if the ItemType is story.
// Validates and sets each field and returns a new story item
// NOTE: used value receiver as we are not mutating
func (cnv ItemConverter) Convert(rank int, item *RawItem) (*Story, error) {
	expectedType := "story"
	if item.ItemType != expectedType {
		return nil, &InvalidItemTypeErr{expectedType, item.ItemType}
	}

	if rank <= 0 {
		return nil, &MinValErr{min: 1, actual: rank}
	}

	validTitle, err := cnv.ValidateStr(item.Title)
	if err != nil {
		return nil, err
	}

	validAuthor, err := cnv.ValidateStr(item.By)
	if err != nil {
		return nil, err
	}

	if !isValidURLScheme(item.URL) {
		return nil, &InvalidURLErr{item.URL}
	}

	points, err := cnv.calculatePoints(item)
	if err != nil {
		return nil, err
	}
	comments, err := cnv.countComments(item)
	if err != nil {
		return nil, err
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
