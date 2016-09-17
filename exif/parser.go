package exif

import (
	"fmt"
	"io"
)

type ExifMap map[string]string

// Parse is a package level function which takes an io.Reader
// and parses the contents to return an ExifMap and an error.
func Parse(r io.Reader) (ExifMap, error) {
	var m = make(ExifMap)

	// look for the JPEG header FFD8 which signifies
	// the start of a JPEG file
	header := make([]byte, 2)
	_, err := r.Read(header)
	if err != nil {
		return m, fmt.Errorf("Unable to parse image header: %s", err)
	}

	// An if statement might make sense here but I have used a
	// switch statement in anticipation of further image types.
	switch string(header) {
	case "\xFF\xD8":
		break
	default:
		return m, fmt.Errorf("The image doesn't appear to be a JPEG. Unknown header found.")
	}

	// Set a dummy value on the ExifMap
	m["status"] = "done"

	return m, nil
}
