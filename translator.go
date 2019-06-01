package main

import (
	"strings"
)

// Error codes for translation errors
const (
	ErrConfusedGopher = 100 // word for translation is shortend
	ErrInvalidWord    = 101
)

// TransError is a transalation error stucture
type TransError struct {
	Err  string
	Code int
}

func (e *TransError) Error() string {
	return e.Err
}

// TranslateWord translates a single word to gophers' language.
func TranslateWord(word string) (string, error) {
	if len(word) == 0 {
		return "", &TransError{Code: ErrInvalidWord, Err: "No word was provided"}
	}

	if strings.ContainsAny(word, "’'") {
		return "", &TransError{Code: ErrConfusedGopher, Err: "Gophers can not understand shortened words"}
	}

	word = strings.ToLower(word)

	var builder strings.Builder
	// handle 2.
	if strings.Index(word, "xr") == 0 {
		builder.WriteString("ge")
		builder.WriteString(word)
		return builder.String(), nil
	}

	vowelIdx := strings.IndexAny(word, "aeiou")
	if vowelIdx == -1 {
		vowelIdx = strings.Index(word, "y")
	}
	// handle 1.
	if vowelIdx == 0 {
		builder.WriteString("g")
		builder.WriteString(word)
		return builder.String(), nil
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
	return builder.String(), nil
}

func extractSign(word string) (string, string) {
	var sign string
	if strings.LastIndexAny(word, ",.?!") == len(word)-1 {
		ln := len(word)
		sign = word[ln-1 : ln]
		word = word[:ln-1]
	}
	return word, sign
}

// TranslateSentence translates a whole sentence in gopher
func TranslateSentence(sentence string) (string, error) {
	english := strings.Split(sentence, " ")
	var gopher []string

	for _, word := range english {
		word, sign := extractSign(word)
		translated, err := TranslateWord(word)
		if e, ok := err.(*TransError); ok {
			if e.Code == ErrConfusedGopher {
				continue
			} else {
				return "", err
			}
		}

		gopher = append(gopher, translated+sign)
	}

	return strings.Join(gopher, " "), nil
}
