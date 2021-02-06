package extractor

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Extraction struct {
	OriginPath      string
	DestinationPath string
}

func (e *Extraction) FindMusic() {
	filepath.Walk(e.OriginPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		filename := info.Name()
		extension := filepath.Ext(filename)

		if validExtension(extension) {
			e.extractSong(path, filename)
		}

		return nil
	})

	fmt.Printf("Done!")
}

func (e *Extraction) extractSong(path, filename string) error {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	destinationFile := e.DestinationPath + filename

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return err
	}

	fmt.Printf("Extracting %s \n.\n", path)
	return nil
}

func validExtension(ext string) bool {
	extensions := []string{
		".mp3",
		".wav",
		".aiff",
		".flac",
		".ogg",
		".aac",
		".wma",
	}

	for _, validExt := range extensions {
		if ext == validExt {
			return true
		}
	}

	return false
}
