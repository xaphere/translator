package main

import "testing"

type testData struct {
	english string
	gopher  string
}

func TestTranslateWord(t *testing.T) {
	pass := []testData{
		{english: "apple", gopher: "gapple"},
		{english: "xray", gopher: "gexray"},
		{english: "chair", gopher: "airchogo"},
		{english: "square", gopher: "aresquogo"},
		{english: "cry", gopher: "ycrogo"},
		{english: "my", gopher: "ymogo"},
		{english: "yellow", gopher: "ellowyogo"},
	}

	fail := []testData{
		{english: "don’t", gopher: ""},
		{english: "", gopher: ""},
		{english: "nth", gopher: ""},
	}

	for _, word := range pass {
		translated, err := TranslateWord(word.english)
		if err != nil {
			t.Error(err)
		}
		if word.gopher != translated {
			t.Errorf("Expected '%s' in gopher, but got '%s'", word.gopher, translated)
		}
	}
	for _, word := range fail {
		_, err := TranslateWord(word.english)
		if err == nil {
			t.Errorf("Expected '%s' to return an error", word.english)
		}

	}
}

func TestTranslateSentence(t *testing.T) {
	td := []struct {
		english string
		gopher  string
	}{
		{
			english: "The quick brown fox jumps, over the lazy dog.",
			gopher:  "ethogo uickqogo ownbrogo oxfogo umpsjogo, gover ethogo azylogo ogdogo.",
		},
		{
			english: "He's good, but not that good.",
			gopher:  "oodgogo, utbogo otnogo atthogo oodgogo.",
		},
	}

	for _, te := range td {
		translated, err := TranslateSentence(te.english)

		if err != nil {
			t.Error(err)
		}

		if te.gopher != translated {
			t.Errorf("Expected '%s' in gopher, but got '%s'", te.gopher, translated)
		}
	}
}
