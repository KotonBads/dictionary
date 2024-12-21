package urbandictionary

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/KotonBads/dictionary/api"
)

func (ub UrbanDictionary) fetchWord(query string) (response, error) {
	client := &http.Client{}
	var result response

	url := fmt.Sprintf(
		"https://unofficialurbandictionaryapi.com/api/search?term=%v&strict=%v&matchCase=%v&limit=%v&page=%v&multiPage=false&", 
		query,
		ub.Strict,
		ub.MatchCase,
		ub.Limit,
		ub.Page,
	)
	resp, err := client.Get(url)
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

func (ub UrbanDictionary) Fetch(word string) (api.Word, error) {
	var w api.Word

	q, err := ub.fetchWord(word)
	if err != nil {
		return w, err
	}
	
	if q.StatusCode != 200 {
		return w, fmt.Errorf("received http %d", q.StatusCode)
	}

	for _, i := range q.Data {
		m := api.Meaning{}
		m.PartOfSpeech = fmt.Sprintf(
			"\033]8;;https://www.urbandictionary.com/author.php?author=%s\033\\%s\033]8;;\033\\\033[0m",
			i.Contributor,
			i.Contributor,
		)
		m.Definition = append(m.Definition, i.Meaning)
		m.Examples = append(m.Examples, strings.Split(i.Example, "\n")...)
		w.Meanings = append(w.Meanings, m)
	}

	return w, nil
}
