package main

import (
	"group-songs/extractor"
	"os"
)

func main() {
	e := extractor.Extraction{
		OriginPath:      os.Args[1],
		DestinationPath: os.Args[2],
	}
	e.FindMusic()
}
