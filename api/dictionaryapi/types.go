package dictionaryapi

type DictionaryApi struct {
}

type response []struct {
	Word     string    `json:"word"`
	Meanings []meaning `json:"meanings"`
	Sources  []string  `json:"sourceUrls"`
}

type meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []definition `json:"definitions"`
	Synonyms     []string     `json:"synonyms"`
	Antonyms     []string     `json:"antonyms"`
}

type definition struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}
