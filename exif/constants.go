package exif

const JpegHeader string = "\xFF\xD8"
const App1Header string = "\xFF\xE1"
const ExifHeader string = "Exif"
const TagMarker string = "\x00\x2A"
const ByteOrderMarker string = "\x4D\x4D"

var TagNames = map[uint]string{
	0x0100: "ImageWidth",
	0x010F: "Make",
	0x0110: "Model",
	0x0112: "Orientation",
	0x011A: "XResolution",
	0x011B: "YResolution",
	0x0128: "ResolutionUnit",
	0x0131: "Software",
	0x0132: "DateTime",
	0x0213: "YCbCrPositioning",

	// Special Tags should probably be handled differently
	0x8769: "Exif IFD",
	0x8825: "GPS IFD",
}
