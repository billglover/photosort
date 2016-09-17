package main

import (
	"fmt"
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
	fmt.Printf("File: %s\n", finfo.Name())
	return nil
}
