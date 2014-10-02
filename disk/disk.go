// Copyright 2014 Dynamic Design. All rights reserved.

package disk

import (
	"net/url"
	"os"
	"path"
	"time"

	"github.com/dynamic-group/storage"
	"github.com/dynamic-group/storage/driver"
)

const (
	DirPerm os.FileMode = 0755
)

type Driver struct{}

func (d *Driver) Open(source string) (driver.Bucket, error) {
	return &Bucket{
		base: source,
	}, nil
}

type Bucket struct {
	base string
}

func (b *Bucket) Create(name string) (driver.File, error) {
	subdir := path.Dir(name)
	// If creating in a subdir
	if subdir != "." {
		if err := os.Mkdir(path.Join(b.base, subdir), os.ModeDir|DirPerm); err != nil {
			return nil, err
		}
	}
	return os.Create(path.Join(b.base, name))
}

func (b *Bucket) Open(name string) (driver.File, error) {
	return os.Open(path.Join(b.base, name))
}

func (b *Bucket) Delete(name string) error {
	return os.Remove(path.Join(b.base, name))
}

func (b *Bucket) URL(path string) (*url.URL, error) {
	panic("unimplemented")
}

func (b *Bucket) SignedURL(path string, expires time.Time) (*url.URL, error) {
	panic("unimplemented")
}

func init() {
	storage.Register("disk", &Driver{})
}
