package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var port = flag.Int("port", 8080, "")

func handleWord(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var english, gopher string
	err = json.Unmarshal(body, &struct {
		Word *string `json:"english-word"`
	}{
		Word: &english,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(english) == 0 {
		http.Error(rw, "No english word provided.", http.StatusBadRequest)
		return
	}
	gopher = TranslateWord(english)

	data, err := json.Marshal(struct {
		Word *string `json:"gopher-word"`
	}{
		Word: &gopher,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleSentence(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var english, gopher string
	err = json.Unmarshal(body, &struct {
		Word *string `json:"english-sentence"`
	}{
		Word: &english,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(english) == 0 {
		http.Error(rw, "No english sentence provided.", http.StatusBadRequest)
		return
	}

	gopher = TranslateSentence(english)

	data, err := json.Marshal(struct {
		Word *string `json:"gopher-sentence"`
	}{
		Word: &gopher,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/word", handleWord)
	http.HandleFunc("/sentence", handleSentence)
	http.HandleFunc("/history”", handleWord)

	log.Println("Start server on ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
