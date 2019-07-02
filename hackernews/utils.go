package hackernews

import "net/url"

// file to hold utility functions

// Returns true if the passed string is a valid url.
// Only allows HTTP or HTTPS
func isValidURLScheme(testURL string) bool {
	_, err := url.ParseRequestURI(testURL)
	if err != nil {
		return false
	}
	return true
}
