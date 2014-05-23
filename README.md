# Storage

Storage is a Go package that provides a single interface for Amazon S3 and Google Cloud Storage through the use of drivers that wrap the well known client libraries for these cloud storage services. There is also a disk-based driver for much simpler needs.

## Installation

Install Storage using the go get command:

    $ go get github.com/dynamic-design/storage

Install one (or more) adapters:

    # Amazon S3
    $ github.com/dynamic-design/storage/s3

    # Google Cloud Storage
    $ github.com/dynamic-design/storage/gcs

    # Local Disk
    $ github.com/dynamic-design/storage/disk

## Documentation

- [Reference](http://godoc.org/github.com/dynamic-design/storage)

## Contributing

Contributions are welcome.

## License

Storage is available under the [BSD (3-Clause) License](http://opensource.org/licenses/BSD-3-Clause).
