package hackernews

import (
	"log"
	"testing"
)

func getTestStory() *Story {

	var testStory = &Story{
		ID:       0,
		Title:    "Mistakes we made adopting event sourcing and how we recovered",
		URL:      "http://natpryce.com/articles/000819.html",
		Author:   "AREALLYFAMOUSAUTHOR",
		Points:   110,
		Comments: 14,
		Rank:     0,
	}
	return testStory

}

func TestValidateStr_empty(t *testing.T) {

	log.Println("Testing storyopts validate empty string options string")

	storyOpts := &StoryOpts{false, false, 0, 0, 0}

	commentTests := []struct {
		input       string
		expected    string
		expectedErr error
	}{
		{"not empty", "not empty", nil},
		{"", "", EmptyStringErr},
	}
	for _, test := range commentTests {
		t.Run(test.input, func(t *testing.T) {
			test := test //capture range variable
			t.Parallel()
			validatedStr, err := storyOpts.ValidateStr(test.input)
			if validatedStr != test.expected {
				t.Errorf("Validated string was incorrect. \n\t Expected %s got %s", test.expected, validatedStr)
			}
			if err != test.expectedErr {
				t.Errorf("Expected error was incorrect. \n\t Expected %s got %s", test.expectedErr.Error(), err.Error())
			}
		})
	}
}

func TestValidateStr_maxStringLength(t *testing.T) {
	log.Println("Testing storyopts validate string maxlength opts")

	storyOpts := &StoryOpts{false, true, 0, 0, 0}
	if _, err := storyOpts.ValidateStr("random"); err != MaxStringErr {
		t.Errorf("Expected error was incorrect. \n\t Expected %s got %s", MaxStringErr, err.Error())
	}

	storyOpts = &StoryOpts{true, true, 20, 0, 0}

	commentTests := []struct {
		input       string
		expected    string
		expectedErr error
	}{
		{"a string", "a string", nil},
		{"A really long string thats too long", "A really long string", nil},
		{"", "", nil}, //testing with empty string allowed!
	}
	for _, test := range commentTests {
		t.Run(test.input, func(t *testing.T) {
			test := test //capture range variable
			t.Parallel()
			validatedStr, err := storyOpts.ValidateStr(test.input)
			if validatedStr != test.expected {
				t.Errorf("Validated string was incorrect. \n\t Expected %s got %s", test.expected, validatedStr)
			}
			if err != test.expectedErr {
				t.Errorf("Expected error was incorrect. \n\t Expected %s got %s", test.expectedErr.Error(), err.Error())
			}
		})
	}
}
