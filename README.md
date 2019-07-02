# HN-Scraper

A hackernews scraper implemented in Go that retrieves the top posts from hackernews and prints them to stdout as json.


# Built With 
* Go     1.11.5 
* linux  18.10
* Docker 18.09.7

# Libraries Used

This project has been built with trying to keep dependencies to a minimum

* go-cmp - Used to help in testing for deep equality


# Instructions to run

Before running ensure you are in the root folder

A dockerfile for this project has been provided, so commands can be executed using go or docker.

## Install

### Docker

  First build the docker image

    ```
    docker build -t hn-scraper .
    ```

  To run 

    ```
    docker run hn-scraper --posts n
    ```

  where n is how many posts you want to scrape

#### Golang


  First build an executable

    ```
    go build
    ```

  To run 

    ```
    hn-scraper --posts n
    ```
    
  where n is how many posts you want to scrape


## Test

Before running tests ensure you are in the root folder

A dockerfile for testing this project has been provided, so tests can be executed using go or docker.



### Docker

  First build the docker image

    ```
    docker build -t hn-scraper-tests .
    ```

  To run tests

    ```
    docker run hn-scraper-tests
    ```


### Golang

  Run the following command 

    ```
    go tests ./...
    ```

  Note: ./... will ensure all tests in packages are run.


# Notes

### Structure

Most processing happens in the hackernews package.

There are 4 main files.
    * api.go handles networks calls to the hackernews api
    * item.go contains struct definitions for different types of items from hackernews
    * itemOpts.go contains struct definitions for story options, Which defines customizable constraints for creating a story. For example, you can control if there is a maxlength enforced.
    * Errors.go contains definitions for errors 
    * utils.go contains general helper functions 

### Tests

In order to help with tests some data from the hackernews api has been saved in json files in the testdata folder.
These are loaded and tested against.
This was done to keep testing files cleaner, since there is quite a lot of data in the json files and also it makes it easier to add more test data in the future.
We could simply add another json file for an item for example.

Some more testing could've been added. 
For example testing the correct errors are returned when expected or some smaller and simpler functions haven't been tested to save time.
Also tests could be cleaned up slightly with more helper functions.

### printing

Currently in the main.go stories may be received and printed out of order.
In order to prevent this, we can insert each story into a slice, using the rank as the index to insert.
but it is not really an issue so ignored for now.

Also if a story is invalid. For example has an invalid uri, then it will print an error in its place, maybe it should just print nothing?
We could improve this by making it get additional stories until it meets the required number.
This would simply involve keeping track of the ids (and possibly rank/index) of retrieved stories and then using the following ids from the list of story ids that we have. Since the limit is 100 and we always get the top 500. For example, if we ask for 50, and 2 of them are invalid. Then we can get story 51 and 52 as we have their ids aswell.

