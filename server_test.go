package main

import (
	"encoding/json"
	"testing"
)

func TestWordHandle(t *testing.T) {
	output, sErr := handleWord(json.RawMessage(`{"english-word":""}`))
	if sErr == nil {
		t.Fatal("Expected error when no word is provided")
	}

	output, sErr = handleWord(json.RawMessage(`{"english":"hello"}`))
	if sErr == nil {
		t.Fatal("Expected error when wrong json is provided")
	}

	output, sErr = handleWord(json.RawMessage(`{"english-word":"hello"}`))
	if sErr != nil {
		t.Error(sErr)
	}

	var gopher string
	err := json.Unmarshal(output, &struct {
		Word string `json:"gopher-word"`
	}{Word: gopher})

	if err != nil {
		t.Error(err)
	}
}

func TestSentenceHandle(t *testing.T) {
	output, sErr := handleSentence(json.RawMessage(`{"english-sentence":""}`))
	if sErr == nil {
		t.Fatal("Expected error when no sentence is provided")
	}

	output, sErr = handleSentence(json.RawMessage(`{"english":"hello world."}`))
	if sErr == nil {
		t.Fatal("Expected error when wrong json is provided")
	}

	output, sErr = handleSentence(json.RawMessage(`{"english-sentence":"hello world."}`))
	if sErr != nil {
		t.Error(sErr)
	}

	var gopher string
	err := json.Unmarshal(output, &struct {
		Word string `json:"gopher-sentence"`
	}{Word: gopher})

	if err != nil {
		t.Error(err)
	}

}
