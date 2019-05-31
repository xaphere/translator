package main

import "testing"

func TestTranslateWord(t *testing.T) {
	words := []struct {
		english string
		gopher  string
	}{
		{english: "apple", gopher: "gapple"},
		{english: "xray", gopher: "gexray"},
		{english: "chair", gopher: "airchogo"},
		{english: "square", gopher: "aresquogo"},
		{english: "don’t", gopher: ""},
		{english: "cry", gopher: "ycrogo"},
		{english: "my", gopher: "ymogo"},
		{english: "yellow", gopher: "ellowyogo"},
		{english: "nth", gopher: "nthogo"},
	}

	for _, word := range words {
		translated := TranslateWord(word.english)

		if word.gopher != translated {
			t.Errorf("Expected %s in gopher, but got %s", word.gopher, translated)
		}
	}

}
