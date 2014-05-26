// Copyright 2014 Dynamic Design. All rights reserved.

package s3

import (
	"bytes"
	"io"
	"mime"
	"os"
	"path"
	"time"

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
		return f.b.PutReader(f.name, &f.buf, int64(f.buf.Len()), mime.TypeByExtension(f.name), s3.Private)
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
	resp, err := f.b.List(f.name, "/", "", 0)
	if err != nil {
		return nil, err
	}

	for _, key := range resp.Contents {
		fi = append(fi, &FileInfo{
			name: key.Key,
			dir:  false,
			size: key.Size,
		})
	}

	return fi, nil
}

func (f *File) Stat() (fi os.FileInfo, err error) {
	panic("unimplemented")
}

func (f *File) Write(b []byte) (n int, err error) {
	return f.buf.Write(b)
}

type FileInfo struct {
	name    string
	dir     bool
	size    int64
	modTime time.Time
}

func (fi *FileInfo) Name() string {
	return path.Base(fi.name)
}

func (fi *FileInfo) Size() int64 {
	return fi.size
}

func (fi *FileInfo) Mode() os.FileMode {
	panic("unimplemented")
}

func (fi *FileInfo) ModTime() time.Time {
	return fi.modTime
}

func (fi *FileInfo) IsDir() bool {
	return fi.dir
}

func (fi *FileInfo) Sys() interface{} {
	return nil
}
