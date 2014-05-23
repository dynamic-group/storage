// Copyright 2014 Dynamic Design. All rights reserved.

package disk

import (
	"os"
	"path"

	"github.com/dynamic-design/storage"
	"github.com/dynamic-design/storage/driver"
)

type Driver struct{}

func (d *Driver) Open(source string) (driver.Bucket, error) {
	return &Bucket{source}, nil
}

type Bucket struct {
	base string
}

func (b *Bucket) Create(name string) (driver.File, error) {
	return os.Create(path.Join(b.base, name))
}

func (b *Bucket) Open(name string) (driver.File, error) {
	return os.Open(path.Join(b.base, name))
}

func init() {
	storage.Register("disk", &Driver{})
}
