package urbandictionary

type UrbanDictionary struct {
	Strict bool
	Limit int
	MatchCase bool
	Page int
}

type response struct {
	StatusCode int    `json:"statusCode"`
	Term       string `json:"term"`
	Found      bool   `json:"found"`
	Params     params `json:"params"`
	TotalPages string `json:"totalPages"`
	Data       []data `json:"data"`
}

type params struct {
	Strict     string `json:"strict"`
	Limit      string `json:"limit"`
	MatchCase  string `json:"matchCase"`
	ScrapeType string `json:"scrapeType"`
	Page       string `json:"page"`
	MultiPage  string `json:"multiPage"`
}

type data struct {
	Word        string `json:"word"`
	Meaning     string `json:"meaning"`
	Example     string `json:"example"`
	Contributor string `json:"contributor"`
	Date        string `json:"date"`
}