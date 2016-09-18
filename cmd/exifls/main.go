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

	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	em, err := exif.Parse(f)
	if err != nil {
		return err
	}

	for i, tag := range em {
		log.Printf("\t%s : %s\n", i, tag)
	}

	return nil
}
