package main

import (
	"encoding/json"
	"testing"
)

func TestHistoryStore(t *testing.T) {

	var history History

	english := "en-test"
	gopher := "go-test"
	gopher2 := "go-test2"

	history.Store(english, gopher)
	gf, ok := history.Load(english)
	if !ok {
		t.Error("Exected stored translation to persist in history")
	}

	if gf != gopher {
		t.Error("Expected gopher translation to persist in history")
	}

	history.Store(english, gopher2)
	gf, ok = history.Load(english)

	if gf != gopher2 {
		t.Error("Expected new Stores to override history")
	}

	data := history.GetData()
	word, ok := data[english]
	if !ok || word != gopher2 {
		t.Error("Expected history data to contain every stored word pair ")
	}
}

func TestHistoryToJSON(t *testing.T) {

	var history History

	data := map[string]string{"dfg": "dfg", "abcd": "abcd", "abc": "abc"}

	for key, val := range data {
		history.Store(key, val)
	}

	jsonData, err := json.Marshal(&history)
	if err != nil {
		t.Error(err)
	}

	var pairs []entry
	err = json.Unmarshal(jsonData, &struct {
		History []entry `json:"history"`
	}{History: pairs})
	if err != nil {
		t.Error(err)
	}
}
