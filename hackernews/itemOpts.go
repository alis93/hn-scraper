package hackernews

import "fmt"

// Options for creating a story.
type StoryOpts struct {
	emptyStringsAllowed    bool
	enforceMaxStringLength bool
	maxStringLength        int
	minComments            int
	minPoints              int
}

func NewStoryOpts(emptyStringAllowed, enforceMaxStrLength bool, maxStrLength, minComments, minPoints int) (*StoryOpts, error) {
	if enforceMaxStrLength && maxStrLength <= 0 {
		return nil, fmt.Errorf("If enforceMaxStringLength is set, then maxStringLength must be more than 0")
	}
	if minComments <= 0 {
		return nil, &MinValErr{1, minComments}
	}
	if minPoints <= 0 {
		return nil, &MinValErr{1, minPoints}
	}
	return &StoryOpts{
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
func (opts StoryOpts) ValidateStr(str string) (string, error) {
	if !opts.emptyStringsAllowed && len(str) <= 0 {
		return "", EmptyStringErr
	}
	if opts.enforceMaxStringLength && len(str) > opts.maxStringLength {
		if opts.maxStringLength <= 0 {
			return "", MaxStringErr
		}
		return str[:opts.maxStringLength], nil
	}
	return str, nil
}

// // Tests if number of comments is more than minimum
// func (opts StoryOpts) hasValidComments(story *Story) bool {
// 	return story.Comments > opts.minComments
// }

// // Tests if number of points is more than minimum
// func (opts StoryOpts) hasValidPoints(story *Story) bool {
// 	return story.Points > opts.minPoints
// }

// Tests and validates author string.
// Returns error if invalid.
// Returns string if valid.
// truncates to required length if enforceMaxStringLength and maxStringLength are set
// func (opts StoryOpts) ValidateAuthor(story *Story) (string, error) {
// 	return opts.validateStr(story.Author)
// }

// // Same as Validate Author but for title
// func (opts StoryOpts) ValidateTitle(story Item) (string, error) {
// 	return opts.validateStr(story.Title)
// }
