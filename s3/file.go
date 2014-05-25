// Copyright 2014 Dynamic Design. All rights reserved.

package s3

import (
	"bytes"
	"io"
	"os"

	"launchpad.net/goamz/s3"
)

type File struct {
	b *Bucket

	name string
	mode int

	rc  io.ReadCloser
	buf bytes.Buffer
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Close() error {
	if f.mode == 0 {
		return f.rc.Close()
	} else {
		return f.b.PutReader(f.name, &f.buf, int64(f.buf.Len()), "", s3.PublicReadWrite)
	}
}

func (f *File) Read(b []byte) (n int, err error) {
	if f.rc == nil {
		rc, err := f.b.Bucket.GetReader(f.name)
		if err != nil {
			return 0, err
		}

		f.rc = rc
	}

	return f.rc.Read(b)
}

func (f *File) Readdir(n int) (fi []os.FileInfo, err error) {
	panic("unimplemented")
}

func (f *File) Stat() (fi os.FileInfo, err error) {
	panic("unimplemented")
}

func (f *File) Write(b []byte) (n int, err error) {
	return f.buf.Write(b)
}
