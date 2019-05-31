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

type translation struct {
	English string
	Gopher  string
}

func (t *translation) UnmarshalJSON(b []byte) error {
	in := struct {
		Word string `json:"english-word"`
	}{}
	if err := json.Unmarshal(b, &in); err != nil {
		return err
	}
	t.English = in.Word
	return nil
}

func (t translation) MarshalJSON() ([]byte, error) {
	out := struct {
		Word string `json:"gopher-word"`
	}{Word: t.Gopher}
	return json.Marshal(out)
}

type history []translation

var hr history

func handleWord(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tr translation

	err = json.Unmarshal(body, &tr)
	if err != nil {
		log.Fatal(err)
	}
	tr.Gopher = TranslateWord(tr.English)

	hr = append(hr, tr)

	err = json.NewEncoder(rw).Encode(tr)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	flag.Parse()

	http.HandleFunc("/word", handleWord)
	http.HandleFunc("/sentence", handleWord)
	http.HandleFunc("/history”", handleWord)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
