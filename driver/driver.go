// Copyright 2014 Dynamic Design. All rights reserved.

package driver

import (
	"os"
)

type Driver interface {
	Open(source string) (Bucket, error)
}

type Bucket interface {
	Create(name string) (File, error)
	Open(name string) (File, error)
}

type File interface {
	Name() string
	Close() error
	Read(b []byte) (n int, err error)
	Readdir(n int) (fi []os.FileInfo, err error)
	Stat() (fi os.FileInfo, err error)
	Write(b []byte) (n int, err error)
}
