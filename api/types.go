package api

type Word struct {
	Meanings []Meaning
	Synonyms []string
	Antonyms []string
	Sources  []string
}

type Meaning struct {
	PartOfSpeech string
	Definition   []string
	Examples     []string
}

type DictionaryProvider interface {
	Fetch(string) (Word, error)
}
