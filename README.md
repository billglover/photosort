# photosort
There are two objectives for code found in this repository; first and formost, to facilitate my learning of [Go](https://golang.org), in addition I'm hoping to build a collection of utilities that help organise an ever growing folder of photos pulled off my phone.

## Part 1 - Exif, the Exchangeable image format

Much of the data we use when sorting photos is stored as hidden meta-data within the image files, the most common format for this data is [Exif](https://en.wikipedia.org/wiki/Exif). There are many existing utilities for reading and modifying this data, but in building my own I'm hoping to familiarise myself with using a couple of packages in the standard library, `io` and `bufio`.

The idea for this command line utility, `exifls` is to take a path to a set of images (`.jpg` only for now) and to list the data extracted from the images.

To build and run this utility against a sample set of images use the following commands from the root folder of the project.

    go build -v ./cmd/...
    ./exifls ./sample_data/*

### Approach

There is an article by Ben Johnson (@benbjohnson) on the [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.176t0epef) that has been doing the rounds in the Go community. On first read, I didn't understand many of the things talked about. In building `exifls` I'm going to attempt to follow the package structure outlined in the article. Some other thoughts on my approach are:

 - The file path will be passed as a command line argument
 - Exif data will be listed to `stdout`
 - Errors in parsing one file will not cause the pogram to crash
 - Errors will be written to `stderr`

### Limitations

 - None of this works yet

### References

 - Exchangeable image file format for digital still cameras: Exif Version 2.2 ([pdf](http://www.exif.org/Exif2-2.PDF))
 - [Description of the Exif file format](http://www.media.mit.edu/pia/Research/deepview/exif.html)
 - [How is Exif info encoded](http://stackoverflow.com/questions/1821515/how-is-exif-info-encoded)
