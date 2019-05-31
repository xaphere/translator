package main

import (
	"strings"
)

// TranslateWord translates a single word to gophers' language
func TranslateWord(word string) string {

	// don't confuse gophers with shortened words
	if strings.ContainsAny(word, "’'") {
		return ""
	}

	var builder strings.Builder
	// handle 2.
	if strings.Index(word, "xr") == 0 {
		builder.WriteString("ge")
		builder.WriteString(word)
		return builder.String()
	}

	vowelIdx := strings.IndexAny(word, "aeiou")
	if vowelIdx == -1 {
		vowelIdx = strings.Index(word, "y")
	}
	// handle 1.
	if vowelIdx == 0 {
		builder.WriteString("g")
		builder.WriteString(word)
		return builder.String()
	}

	// handle 4.
	if vowelIdx >= 2 && word[vowelIdx-1:vowelIdx+1] == "qu" {
		vowelIdx++
	}

	// handle 3.
	if vowelIdx == -1 {
		// set to 0 to handle voweless words
		vowelIdx = 0
	}

	builder.WriteString(word[vowelIdx:len(word)])
	builder.WriteString(word[0:vowelIdx])
	builder.WriteString("ogo")

	return builder.String()
}
