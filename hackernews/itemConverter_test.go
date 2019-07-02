package hackernews

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func loadItem(t *testing.T, itemId int) *RawItem {

	file_template := "item_%d.json"
	itemBytes := helperLoadBytes(t, fmt.Sprintf(file_template, itemId))
	rawItem := &RawItem{}
	if err := json.Unmarshal(itemBytes, rawItem); err != nil {
		t.Fatal(err)
	}
	return rawItem
}

var storyTests = []struct {
	id            int
	expectedStory *Story
}{
	{20324021, &Story{ID: 0, Title: "Mistakes we made adopting event sourcing and how we recovered", URL: "http://natpryce.com/articles/000819.html", Author: "moks", Points: 110, Comments: 14, Rank: 1}},
	{20325395, &Story{ID: 0, Title: "Google’s robots.txt parser is now open source", URL: "https://opensource.googleblog.com/2019/07/googles-robotstxt-parser-is-now-open.html", Author: "dankohn1", Points: 570, Comments: 147, Rank: 2}},
}

func TestValidateStrEmptyStringAllowed(t *testing.T) {

	log.Println("Testing item converter validate empty string option")

	cnv := &ItemConverter{false, false, 0, 0, 0}

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
			validatedStr, err := cnv.ValidateStr(test.input)
			if validatedStr != test.expected {
				t.Errorf("Validated string was incorrect. \n\t Expected %s got %s", test.expected, validatedStr)
			}
			if err != test.expectedErr {
				t.Errorf("Expected error was incorrect. \n\t Expected %s got %s", test.expectedErr.Error(), err.Error())
			}
		})
	}
}

func TestValidateStrMaxStringLength(t *testing.T) {
	log.Println("Testing item converter validate string maxlength option")

	cnv := &ItemConverter{false, true, 0, 0, 0}
	if _, err := cnv.ValidateStr("random"); err != MaxStringErr {
		t.Errorf("Expected error was incorrect. \n\t Expected %s got %s", MaxStringErr, err.Error())
	}

	cnv = &ItemConverter{true, true, 20, 0, 0}

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
			validatedStr, err := cnv.ValidateStr(test.input)
			if validatedStr != test.expected {
				t.Errorf("Validated string was incorrect. \n\t Expected %s got %s", test.expected, validatedStr)
			}
			if err != test.expectedErr {
				t.Errorf("Expected error was incorrect. \n\t Expected %s got %s", test.expectedErr.Error(), err.Error())
			}
		})
	}
}

func TestCalculatePoints(t *testing.T) {
	log.Println("Testing Calculate points in item")
	cnv := &ItemConverter{false, false, 0, 0, 0}

	for _, test := range storyTests {
		t.Run(strconv.Itoa(test.id), func(t *testing.T) {
			test := test //capture range variable
			t.Parallel()
			item := loadItem(t, test.id)
			points := cnv.calculatePoints(item)
			if points != test.expectedStory.Points {
				t.Errorf("Points incorrect. \n\t Expected %d Actual : %d", test.expectedStory.Points, points)
			}
		})
	}
}

func TestCountComments(t *testing.T) {
	cnv := &ItemConverter{false, false, 0, 0, 0}
	log.Println("Testing count comments in item")
	for _, test := range storyTests {
		t.Run(strconv.Itoa(test.id), func(t *testing.T) {
			test := test //capture range variable
			t.Parallel()
			item := loadItem(t, test.id)
			comments := cnv.countComments(item)
			if comments != test.expectedStory.Comments {
				t.Errorf("Comments incorrect. \n\t Expected %d Actual : %d", test.expectedStory.Comments, comments)
			}
		})
	}
}

func TestConvertToStory_no_options(t *testing.T) {
	log.Println("Testing conversion of raw item to story, without options")
	for idx, test := range storyTests {
		t.Run(strconv.Itoa(test.id), func(t *testing.T) {
			test := test //capture range variable
			t.Parallel()
			item := loadItem(t, test.id)
			cnv := &ItemConverter{false, false, 0, 0, 0}
			story, err := cnv.ConvertToStory(idx+1, item)

			if err != nil {
				t.Fatalf("Failed to convert story: Reason %s", err.Error())
			}

			if story.Rank < 0 {
				t.Errorf("Rank cannot be less than 0. \t Rank was %d", story.Rank)

			}

			if story.Comments < 0 {
				t.Errorf("number of comments cannot be less than 0. \t points was %d", story.Comments)
			}

			if story.Points < 0 {
				t.Errorf("points cannot be less than 0. \t points was %d", story.Points)
			}

			if cmp.Equal(story, test.expectedStory) {
				t.Errorf("Convert to story was wrong. \n\t Expected %+v Actual : %+v", test.expectedStory, story)
			}
		})
	}
}

// We could create a table of options to try for each test
// This would make it easier to combine with previous test, since no_opts will be a test case in that table
func TestConvertToStory_with_options(t *testing.T) {
	log.Println("Testing conversion of raw item to story, with options set")
	for idx, test := range storyTests {

		t.Run(strconv.Itoa(test.id), func(t *testing.T) {
			test := test //capture range variable
			t.Parallel()

			item := loadItem(t, test.id)
			stringLength := 256

			cnv := &ItemConverter{false, true, stringLength, 0, 0}

			story, err := cnv.ConvertToStory(idx+1, item)

			if err != nil {
				t.Fatalf("Failed to convert story: Reason %s", err.Error())
			}

			if story.Title == "" {
				t.Errorf("Title cannot be an empty string. \n\t Expected %s Actual %s", test.expectedStory.Title, story.Title)
			}

			if len(story.Title) > stringLength {
				t.Errorf("Length of title was more than max stringlength. \n\t max stringlength was %d Actual %d", stringLength, len(story.Title))
			}

			if story.Author == "" {
				t.Errorf("Author cannot be an empty string. \n\t Expected %s Actual %s", test.expectedStory.Author, story.Author)
			}

			if len(story.Author) > stringLength {
				t.Errorf("Length of Author was more than max stringlength. \n\t max stringlength was %d Actual %d", stringLength, len(story.Title))
			}

			if cmp.Equal(story, test.expectedStory) {
				// if story != test.expectedStory {
				t.Errorf("Convert to story was wrong. \n\t Expected %+v Actual : %+v", test.expectedStory, story)
			}
		})
	}
}
