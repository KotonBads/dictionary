package dictionaryapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/KotonBads/dictionary/api"
)

func fetchWord(query string) (response, error) {
	client := &http.Client{}
	var result response

	resp, err := client.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + query)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(out, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (d DictionaryApi) Fetch(word string) (api.Word, error) {
	var w api.Word

	q, err := fetchWord(word)
	if err != nil {
		return w, err
	}

	for _, i := range q {
		for _, j := range i.Meanings {
			for _, k := range j.Definitions {
				m := api.Meaning{}
				m.PartOfSpeech = j.PartOfSpeech
				m.Definition = append(m.Definition, k.Definition)
				m.Examples = append(m.Examples, k.Example)
				w.Meanings = append(w.Meanings, m)
				w.Antonyms = append(w.Antonyms, k.Antonyms...)
				w.Synonyms = append(w.Synonyms, k.Synonyms...)
			}
			w.Antonyms = append(w.Antonyms, j.Antonyms...)
			w.Synonyms = append(w.Synonyms, j.Synonyms...)
		}
		w.Sources = append(w.Sources, i.Sources...)
	}

	return w, nil
}
