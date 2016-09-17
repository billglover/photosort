package exif

import (
	"bufio"
	"encoding/binary"
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
	case JpegHeader:
		break
	default:
		return m, fmt.Errorf("The image doesn't appear to be a JPEG. Unknown header found.")
	}

	// look for the APP1 marker which signifies the start
	// of the EXIF data area
	b := make([]byte, 2)
	for string(b) != App1Header {
		_, err := r.Read(b)
		if err != nil {
			return m, fmt.Errorf("Unable to locate the APP1 header: %s", err)
		}
	}

	// the next four bytes indicate the size of the APP1
	// data area
	size := make([]byte, 2)
	_, err = r.Read(size)
	if err != nil {
		return m, fmt.Errorf("Unable to determine the size of the APP1 header: %s", err)
	}

	// look for the EXIF header
	exifHeader := make([]byte, 4)
	_, err = r.Read(exifHeader)
	if err != nil {
		return m, fmt.Errorf("Unable to locate the Exif header: %s", err)
	}

	if string(exifHeader) != ExifHeader {
		return m, fmt.Errorf("Unable to parse the Exif header. Found '%s' expected 'Exif'", string(exifHeader))
	}

	// look for the padding following the EXIF header
	p := make([]byte, 2)
	_, err = r.Read(p)
	if err != nil {
		return m, fmt.Errorf("Unable to parse the Exif header: %s", err)
	}

	// look for the byte order marker
	// for JPEG files this should always be MM or 0x4949
	bom := make([]byte, 2)
	_, err = r.Read(bom)
	if err != nil {
		return m, fmt.Errorf("Unable to locate the Exif byte order marker: %s", err)
	}

	if string(bom) != ByteOrderMarker {
		return m, fmt.Errorf("Unexpected byte order marker for JPEG file. Found '%s' expected 'MM'", string(bom))
	}

	// look for the TAG marker
	// for JPEG files this should always be 2A00
	tag := make([]byte, 2)
	_, err = r.Read(tag)
	if err != nil {
		return m, fmt.Errorf("Unable to locate the TAG marker: %s", err)
	}

	if string(tag) != TagMarker {
		return m, fmt.Errorf("Unexpected TAG marker for JPEG file. Found '%s' expected '%s'", string(tag), TagMarker)
	}

	// look for the offset to the first IFD
	// this is an offset from the header so it includes
	// 2 bytes for the BOM
	// 2 bytes for the TAG marker
	// 4 bytes for the offset
	o := make([]byte, 4)
	_, err = r.Read(o)
	if err != nil {
		return m, fmt.Errorf("Unable to locate the offset to the first IFD: %s", err)
	}

	offset := int(binary.BigEndian.Uint32(o))
	offset = offset - 8

	// at this point I'm not sure I can discard or skip
	// bytes without using a buffio.Reader
	buff := bufio.NewReader(r)
	buff.Discard(offset)

	entries := make([]byte, 2)
	_, err = buff.Read(entries)
	if err != nil {
		return m, fmt.Errorf("Unable to read the number of entries in the IFD: %s", err)
	}
	tags := int(binary.BigEndian.Uint16(entries))

	for tags > 0 {

		// read the tag ID
		tag := make([]byte, 2)
		_, err = buff.Read(tag)
		if err != nil {
			return m, fmt.Errorf("Unable to parse Exif tag type: %s", err)
		}

		// determine the data format
		format := make([]byte, 2)
		_, err = buff.Read(format)
		if err != nil {
			return m, fmt.Errorf("Unable to parse Exif tag data format: %s", err)
		}

		// determine the number of components in the tag value
		components := make([]byte, 4)
		_, err = buff.Read(components)
		if err != nil {
			return m, fmt.Errorf("Unable to parse Exif tag data format: %s", err)
		}
		//nc := int(binary.BigEndian.Uint32(components))
		name, ok := TagNames[uint(binary.BigEndian.Uint16(tag))]
		if ok == false {
			name = fmt.Sprintf("UnknownTag_%02d", tags)
		}
		m[name] = "unknown"

		// skip the remaining tag bytes
		buff.Discard(4)

		tags = tags - 1
	}

	return m, nil
}
