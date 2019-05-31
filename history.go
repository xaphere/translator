package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Entry struct {
	En string
	Gf string
}

func (e *Entry) MarshalJSON() ([]byte, error) {

	str := fmt.Sprintf("{\"%s\":\"%s\"}", e.En, e.Gf)
	b := []byte(str)
	return json.RawMessage(b), nil
}

type History struct {
	Words []Entry `json:"history"`
}

func (h *History) Remember(en string, gf string) {
	h.Words = append(h.Words, Entry{En: en, Gf: gf})
}

func (h History) Len() int {
	return len(h.Words)
}
func (h History) Less(i, j int) bool {
	return strings.Compare(h.Words[i].En, h.Words[j].En) < 0
}
func (h History) Swap(i, j int) {
	h.Words[i].En, h.Words[j].En = h.Words[j].En, h.Words[i].En
	h.Words[i].Gf, h.Words[j].Gf = h.Words[j].Gf, h.Words[i].Gf
}
