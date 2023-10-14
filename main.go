package main

import (
	"github.com/asticode/go-astisub"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		panic("Usage: blah './path/to/subtitles.src.srt' 'blah blah' './path/to/subtitles.dst.srt'")
	}

	input := os.Args[1]
	text := os.Args[2]
	output := os.Args[3]

	subtitles, err := astisub.OpenFile(input)

	if err != nil {
		panic(err)
	}

	// Using the index because the value is copied
	for subtitleIndex := range subtitles.Items {
		subtitle := subtitles.Items[subtitleIndex]
		for lineIndex := range subtitle.Lines {
			line := subtitle.Lines[lineIndex]
			for lineItemIndex := range line.Items {
				subtitles.Items[subtitleIndex].Lines[lineIndex].Items[lineItemIndex].Text = text
			}
		}
	}

	err = subtitles.Write(output)
	if err != nil {
		panic(err)
	}
}
