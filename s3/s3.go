// Copyright 2014 Dynamic Design. All rights reserved.

package s3

import (
	// "os"
	// "path"
	"strings"

	"github.com/dynamic-design/storage"
	"github.com/dynamic-design/storage/driver"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

type Driver struct{}

func (d *Driver) Open(source string) (driver.Bucket, error) {
	flds := map[string]string{}
	for _, fld := range strings.Fields(source) {
		toks := strings.SplitN(fld, "=", 2)
		flds[toks[0]] = toks[1]
	}
	c := s3.New(aws.Auth{
		AccessKey: flds["key"],
		SecretKey: flds["secret"],
	}, aws.Regions[flds["region"]])

	return &Bucket{
		Bucket: c.Bucket(flds["bucket"]),
	}, nil
}

type Bucket struct {
	*s3.Bucket
}

func (b *Bucket) Create(name string) (driver.File, error) {
	return &File{
		b:    b,
		name: name,
		mode: 1,
	}, nil
}

func (b *Bucket) Open(name string) (driver.File, error) {
	return &File{
		b:    b,
		name: name,
		mode: 0,
	}, nil
}

func init() {
	storage.Register("s3", &Driver{})
}
