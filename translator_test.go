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
		{english: "", gopher: ""},
	}

	for _, word := range words {
		translated := TranslateWord(word.english)

		if word.gopher != translated {
			t.Errorf("Expected '%s' in gopher, but got '%s'", word.gopher, translated)
		}
	}

}

func TestTranslateSentence(t *testing.T) {
	td := []struct {
		english string
		gopher  string
	}{{
		english: "The quick brown fox jumps, over the lazy dog.",
		gopher:  "ethogo uickqogo ownbrogo oxfogo umpsjogo, gover ethogo azylogo ogdogo.",
	},
		{
			english: "He's good, but not that good.",
			gopher:  "oodgogo, utbogo otnogo atthogo oodgogo.",
		}}

	for _, te := range td {
		translated := TranslateSentence(te.english)

		if te.gopher != translated {
			t.Errorf("Expected '%s' in gopher, but got '%s'", te.gopher, translated)
		}
	}
}
