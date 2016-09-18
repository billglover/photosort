package main

import (
	"github.com/billglover/photosort/exif"
	"log"
	"os"
	"path/filepath"
)

func main() {
	paths := os.Args[1:]
	for _, p := range paths {
		err := filepath.Walk(p, handleFile)
		if err != nil {
			log.Println(err)
		}
	}
}

// handleFile is called once for each file found in the path. It
// returns an error if it is unable to open the file for read.
func handleFile(p string, finfo os.FileInfo, err error) error {
	log.Printf("parsing file: %s\n", finfo.Name())

	f, fileErr := os.Open(p)
	if fileErr != nil {
		return fileErr
	}
	defer f.Close()

	em, exifErr := exif.Parse(f)
	if exifErr != nil {
		return exifErr
	}

	for i, tag := range em {
		log.Printf("\t%s : %s\n", i, tag)
	}

	return nil
}
