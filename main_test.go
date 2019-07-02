package main

// import (
// 	"os"
// 	"strconv"
// 	"testing"
// )

// func setup() {
// }

// var numPostTests = []struct {
// 	input    int
// 	expected int
// }{
// 	{0, 0},

// 	{10, 10},
// 	// {}
// }

// //The tests
// func TestGetNumPostsArg(t *testing.T) {

// 	for _, test := range numPostTests {
// 		os.Args = []string{"posts", "0"}
// 		test := test // capture range variable for parallelism
// 		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
// 			t.Parallel()
// 			if result := getNumPostsArg(); result != test.expected {
// 				t.Errorf("Got %d want %d", result, test.expected)
// 			}
// 		})
// 	}

// }
