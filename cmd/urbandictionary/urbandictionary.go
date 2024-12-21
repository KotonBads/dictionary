package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/KotonBads/dictionary/api"
	"github.com/KotonBads/dictionary/api/urbandictionary"
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
	exampleStype = baseStyle.
			PaddingLeft(2)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a word to look up.")
		return
	}
	words := os.Args[1:]
	res, err := api.GetDefinition(
		urbandictionary.UrbanDictionary{
			Strict:    false,
			Limit:     10,
			MatchCase: false,
			Page:      1,
		},
		strings.Join(words, "%20"),
	)
	if err != nil {
		panic(err)
	}

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
			exampleStype.Render(
				list.New(
					i.Examples,
				).String(),
			),
		)
		fmt.Println("")
	}
}
