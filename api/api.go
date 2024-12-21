package api

func GetDefinition(provider DictionaryProvider, query string) (Word, error) {
	return provider.Fetch(query)
}