package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/KotonBads/dictionary/api"
	"github.com/KotonBads/dictionary/api/dictionaryapi"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var (
	baseStyle = lipgloss.NewStyle().
			Width(80).
			Foreground(lipgloss.Color("7")).
			TabWidth(2)
	headerStyle = baseStyle.
			Bold(true).
			Foreground(lipgloss.Color("12"))
	textStyle = baseStyle.
			Foreground(lipgloss.Color("10"))
	synonymStyle = baseStyle.
			Foreground(lipgloss.Color("10"))
	antonymStyle = baseStyle.
			Foreground(lipgloss.Color("9"))
	exampleStyle = baseStyle.
			PaddingLeft(2)
	wordStyle = headerStyle.
			Foreground(lipgloss.Color("15"))
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a word to look up.")
		return
	}
	words := os.Args[1:]
	res, err := api.GetDefinition(
		dictionaryapi.DictionaryApi{},
		strings.Join(words, "%20"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(wordStyle.Render(res.Word))

	for _, i := range res.Meanings {
		fmt.Println(headerStyle.Render(i.PartOfSpeech))
		fmt.Println(
			textStyle.Render(
				list.New(
					i.Definition,
				).String(),
			),
		)

		if i.Examples[0] == "" {
			fmt.Println("")
			continue
		}

		fmt.Println(
			exampleStyle.Render(
				list.New(
					i.Examples,
				).String(),
			),
		)
		fmt.Println("")
	}

	fmt.Println(headerStyle.Render("synonyms"))
	fmt.Println(synonymStyle.Render(
		list.New(
			res.Synonyms,
		).String(),
	))
	fmt.Println("")

	fmt.Println(headerStyle.Render("antonyms"))
	fmt.Println(antonymStyle.Render(
		list.New(
			res.Antonyms,
		).String(),
	))
	fmt.Println("")
}
