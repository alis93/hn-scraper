package hackernews

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var timeoutTests = []struct {
	input         int
	expectedValid bool
}{
	{10, true},
	{20, true},
	{0, false},
	{-5, false},
}

func TestNewClient(t *testing.T) {
	log.Println("Testing Client Creation")

	for _, test := range timeoutTests {
		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			test := test //capture range variable
			t.Parallel()
			client, err := NewClient(test.input)

			// Expected to be valid. Should have client and no error
			if test.expectedValid && (client == nil || err != nil) {
				t.Errorf("Client was not created. Expected err to be nil but got \n %s", err.Error())
			}
			// Expected to be invalid. Should have err
			if !test.expectedValid && err == nil {
				t.Errorf("Client was expected to be nil, but was not.")
			}

			// Test url set properly
			expectedUrl := "https://hacker-news.firebaseio.com/v0"
			if client != nil && client.apiURL != expectedUrl {
				t.Errorf("Client apiUrl set incorrectly. \n\t Expected : %s actual %s", expectedUrl, "client.apiURL")
			}

		})

	}

}

func TestGetTopStoryIds(t *testing.T) {
	log.Println("Testing Get top story ids")

	testData := helperLoadBytes(t, "topstories.json")
	parsedTestData := []int{}
	json.Unmarshal(testData, &parsedTestData)
	// setup mock client and server
	client, server := mockServerHelper(testData)
	defer server.Close()

	numStories := -1

	// Test that amount argument validated.
	// function should return err if amount<=0 or amount > 500
	if _, err := client.GetTopStoryIds(numStories); err == nil {
		t.Fatalf("Amount argument is not valid. Must be between 1 and 500, but was %d", numStories)
	}
	numStories = 0
	if _, err := client.GetTopStoryIds(numStories); err == nil {
		t.Fatalf("Amount argument is not valid. Must be between 1 and 500, but was %d", numStories)
	}
	numStories = 600
	if _, err := client.GetTopStoryIds(numStories); err == nil {
		t.Fatalf("Amount argument is not valid. Must be between 1 and 500, but was %d", numStories)
	}

	numStories = 100
	topStories, err := client.GetTopStoryIds(numStories)
	if err != nil {
		t.Fatalf("Failed to load top stories . \n Reason : %s", err.Error())
	}

	if len(topStories) > numStories || len(topStories) <= 0 {
		t.Fatalf("number of top stories is incorrect. \n\t Expected to retrieve at most %d, but got %d", numStories, len(topStories))
	}

	numStories = 3
	topStories, err = client.GetTopStoryIds(numStories)
	if err != nil {
		t.Fatalf("Failed to load top stories . \n Reason : %s", err.Error())
	}
	if len(topStories) > numStories || len(topStories) <= 0 {
		t.Fatalf("number of top stories is incorrect. \n\t Expected to retrieve at most %d, but got %d", numStories, len(topStories))
	}

	if len(parsedTestData) > numStories {
		parsedTestData = parsedTestData[:numStories]
	}

	if !cmp.Equal(topStories, parsedTestData) {
		t.Fatalf("Expected output is not equal. Expected : %v, \n Actual : %v", parsedTestData, topStories)
	}

}

func TestGetItem(t *testing.T) {
	log.Println("Testing Get Item")

	testIds := []int{20324021, 20325395, 20325925, 20328871, 20329699}
	file_template := "item_%d.json"

	for _, testID := range testIds {
		t.Run(strconv.Itoa(testID), func(t *testing.T) {
			testID := testID //capture range variable
			t.Parallel()

			testItem := helperLoadBytes(t, fmt.Sprintf(file_template, testID))

			client, server := mockServerHelper(testItem)
			defer server.Close()

			item, err := client.GetItem(testID)
			if err != nil {
				t.Errorf("Error retrieving item from client. \n Reason: %s", err.Error())
			}

			ts, err := json.Marshal(item)
			if err != nil {
				t.Fatalf("Failed to marshal item into json \n reason :%s", err.Error())
			}

			if bytes.Equal(ts, testItem) {
				t.Fatalf("Expected output is not equal. Expected : %s, \n Actual : %s", testItem, ts)
			}

		})
	}

}
