package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/alis93/hn-scraper/hackernews"
)

var ErrorLog = log.New(os.Stderr,
	"Error: ",
	log.Ldate|log.Ltime|log.Lshortfile)

func main() {

	numPosts := getNumPostsArg()
	fmt.Fprintf(os.Stdout, "Retrieving %d posts\n", numPosts)

	// create hackernews client
	timeout := 5
	client, err := hackernews.NewClient(timeout)
	if err != nil {
		ErrorLog.Fatal(err)
	}

	// get array of top story ids
	storyIds, err := client.GetTopStoryIds(numPosts)
	if err != nil {
		ErrorLog.Fatal(err)
	}
	converter, err := hackernews.NewItemConverter(false, true, 256, 1, 1)
	if err != nil {
		ErrorLog.Fatal(err)
	}

	// Create a channel
	storyChan := make(chan *hackernews.Story)
	var wg sync.WaitGroup
	// loop each id, retrieve item and convert to story then send to channel
	for index, storyId := range storyIds {
		wg.Add(1)
		go func(index, storyId int) {
			rawItem, err := client.GetItem(storyId)
			if err != nil {
				ErrorLog.Printf("Unable to get item with id %d, Reason: %s \n", storyId, err.Error())
			}

			story, err := converter.Convert(index+1, rawItem)
			if err != nil {
				ErrorLog.Printf("Unable to convert item with id %d to story, Reason: %s \n", storyId, err.Error())
			}
			storyChan <- story
			wg.Done()
		}(index, storyId)
	}

	go func() {
		wg.Wait()
		close(storyChan)
	}()

	// listen and wait on channel
	for story := range storyChan {
		// print each story received on channel
		fmt.Println(story)
	}
}

//
func getNumPostsArg() int {
	numPosts := flag.Int("posts", 0, "How many posts to retrieve")
	flag.Parse()
	if *numPosts <= 0 || *numPosts > 100 {
		ErrorLog.Printf("You requested %d posts. Must be between 0 and 100", *numPosts)
		os.Exit(1)
	}
	return *numPosts
}
