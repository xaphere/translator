package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handleWord(rw http.ResponseWriter, req *http.Request) *serverError {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	var english, gopher string
	err = json.Unmarshal(body, &struct {
		Word *string `json:"english-word"`
	}{
		Word: &english,
	})
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	if len(english) == 0 {
		return &serverError{Error: nil, Msg: "No english word provided.", Code: http.StatusBadRequest}
	}

	var ok bool
	if gopher, ok = history.Load(english); !ok {
		gopher, err = TranslateWord(english)
		if err != nil {
			return &serverError{Error: err, Msg: "Translation failed", Code: http.StatusBadRequest}
		}
		history.Store(english, gopher)
	}

	data, err := json.Marshal(struct {
		Word *string `json:"gopher-word"`
	}{
		Word: &gopher,
	})
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	_, err = rw.Write(data)
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	return nil
}

func handleSentence(rw http.ResponseWriter, req *http.Request) *serverError {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	var english, gopher string
	err = json.Unmarshal(body, &struct {
		Word *string `json:"english-sentence"`
	}{
		Word: &english,
	})
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	if len(english) == 0 {
		return &serverError{Error: nil, Msg: "No english sentence provided.", Code: http.StatusBadRequest}
	}

	if gopher, ok := history.Load(english); !ok {
		gopher, err = TranslateSentence(english)
		if err != nil {
			return &serverError{Error: err, Msg: "Translation failed", Code: http.StatusBadRequest}
		}
		history.Store(english, gopher)
	}

	data, err := json.Marshal(struct {
		Word *string `json:"gopher-sentence"`
	}{
		Word: &gopher,
	})
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	_, err = rw.Write(data)
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}
	return nil
}

var history History

func handleHistory(rw http.ResponseWriter, req *http.Request) *serverError {

	data, err := history.ToJSON()
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}

	}
	_, err = rw.Write(data)
	if err != nil {
		return &serverError{Error: err, Msg: err.Error(), Code: http.StatusInternalServerError}
	}
	return nil
}

type serverError struct {
	Error error
	Msg   string
	Code  int
}

type serverHandler struct {
	Handle func(w http.ResponseWriter, r *http.Request) *serverError
	Method string
}

func (sh serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != sh.Method {
		http.Error(w, fmt.Sprintf("%s method not supported", r.Method), http.StatusMethodNotAllowed)
		return
	}

	if err := sh.Handle(w, r); err != nil {
		http.Error(w, err.Msg, err.Code)
	}
}

func initServer(port int) error {
	http.Handle("/word", serverHandler{Handle: handleWord, Method: "POST"})
	http.Handle("/sentence", serverHandler{Handle: handleSentence, Method: "POST"})
	http.Handle("/history", serverHandler{Handle: handleHistory, Method: "GET"})

	log.Println("Start server on ", port)
	strPort := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(strPort, nil)
	if err != nil {
		return err
	}
	return nil
}
