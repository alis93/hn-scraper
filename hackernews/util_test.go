package hackernews

import (
	"testing"
)

var urlTests = []struct {
	input    string
	expected bool
}{
	{"http://google.com", true},
	{"https://bbc.com", true},
	{"https://bbc.com/dklfjsdf", true},
	{"https://www.google.com/search?q=fdjskfd&oq", true},
	{"mongodb://mongodb0.example.com:27017/admin", true},
	{"", false},
	{"Hello how are you", false},
}

func TestIsValidURLScheme(t *testing.T) {

	for _, test := range urlTests {
		t.Run(test.input, func(t *testing.T) {
			test := test //capture range variable
			t.Parallel()
			if isValid := isValidURLScheme(test.input); isValid != test.expected {
				t.Errorf("%s failed. \n\t Expected: %t Actual %t ", test.input, test.expected, isValid)
			}
		})

	}

}
