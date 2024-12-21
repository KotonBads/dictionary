package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/KotonBads/dictionary/api"
	"github.com/KotonBads/dictionary/api/urbandictionary"
	"github.com/KotonBads/dictionary/internal/styles"
	"github.com/charmbracelet/lipgloss/list"
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

	fmt.Println(styles.WordStyle.Render(res.Word))

	for _, i := range res.Meanings {
		fmt.Println(styles.HeaderStyle.Render(i.PartOfSpeech))
		fmt.Println(
			styles.TextStyle.Render(
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
			styles.ExampleStyle.Render(
				list.New(
					i.Examples,
				).String(),
			),
		)
		fmt.Println("")
	}
}
