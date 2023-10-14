package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/asticode/go-astisub"
)

const intFlagCount = 10

func main() {
	input := flag.String("i", "", "input subtitle file")
	output := flag.String("o", "", "output subtitle file")
	text := flag.String("t", "", "text to replace every line with")
	intFlags := parseIntFlags()

	flag.Parse()

	firstIndex, lastIndex := findFirstAndLastLines(intFlags)
	if *input == "" || *output == "" || (*text == "" && firstIndex < 0 && lastIndex < 0) {
		flag.Usage()
		os.Exit(1)
	}

	subtitles, err := astisub.OpenFile(*input)
	if err != nil {
		panic(err)
	}

	// Using the index because the value is copied
	for subtitleIndex := range subtitles.Items {
		subtitle := subtitles.Items[subtitleIndex]

		if *text != "" {
			for lineIndex := range subtitle.Lines {
				line := subtitle.Lines[lineIndex]

				for lineItemIndex := range line.Items {
					subtitles.Items[subtitleIndex].Lines[lineIndex].Items[lineItemIndex].Text = *text
				}
			}
		} else {
			lines := []astisub.Line{}

			for _, flag := range intFlags[firstIndex : lastIndex+1] {
				lines = append(lines, astisub.Line{
					Items: []astisub.LineItem{
						{
							Text:        *flag,
							InlineStyle: subtitle.Lines[0].Items[0].InlineStyle,
							Style:       subtitle.Lines[0].Items[0].Style,
						},
					},
				})
			}

			subtitles.Items[subtitleIndex].Lines = lines
		}
	}

	err = subtitles.Write(*output)
	if err != nil {
		panic(err)
	}
}

func findFirstAndLastLines(list []*string) (int, int) {
	firstIndex := -1
	lastIndex := -1

	for index, value := range list {
		if *value != "" {
			if firstIndex < 0 {
				firstIndex = index
			}

			lastIndex = index
		}
	}

	return firstIndex, lastIndex
}

func parseIntFlags() []*string {
	flags := []*string{}
	for i := 0; i < intFlagCount; i++ {
		flags = append(
			flags,
			flag.String(fmt.Sprintf("%d", i), "", fmt.Sprintf("What to put in line %d", i)),
		)
	}
	return flags
}
