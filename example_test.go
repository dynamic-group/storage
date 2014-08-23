// Copyright 2014 Dynamic Design. All rights reserved.

package storage_test

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/dynamic-design/storage"
	_ "github.com/dynamic-design/storage/disk"
	_ "github.com/dynamic-design/storage/s3"
)

var (
	drivers  map[string]string
	S3key    = flag.String("s3-key", "", "AWS key with s3 access")
	S3secret = flag.String("s3-secret", "", "AWS secret with s3 access")
	S3bucket = flag.String("s3-bucket", "", "S3 bucket name")
	S3region = flag.String("s3-region", "eu-west-1", "S3 AWS region")
)

func TestIntegration(t *testing.T) {

	setup(t)
	defer tearDown()

	for driver, conf := range drivers {
		b, err := storage.Open(driver, conf)
		if err != nil {
			t.Errorf("%s: %s", driver, err)
		}

		path := "subdir/glados.txt"
		content := []byte("You haven't escaped, you know.\n")

		// Create the file
		f, err := b.Create(path)
		if err != nil {
			t.Errorf("%s: %s", driver, err)
		}
		if _, err := f.Write(content); err != nil {
			t.Errorf("%s: %s", driver, err)
		}
		if err := f.Close(); err != nil {
			t.Errorf("%s: %s", driver, err)
		}

		// Open it
		f, err = b.Open(path)
		if err != nil {
			t.Errorf("%s: %s", driver, err)
		}
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, f); err != nil {
			t.Errorf("%s: %s", driver, err)
		}
		if err := f.Close(); err != nil {
			t.Errorf("%s: %s", driver, err)
		}
		if !bytes.Equal(buf.Bytes(), content) {
			t.Errorf("%s: %s", driver, "read file content is not the same as we stored")
		}

		// Delete the file
		if err := b.Delete(path); err != nil {
			t.Errorf("%s: %s", driver, err)
		}
		// tearDown()
	}
}

func setup(t *testing.T) {
	flag.Parse()

	drivers = make(map[string]string)

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	drivers["disk"] = dir

	drivers["s3"] = fmt.Sprintf("key=%s secret=%s bucket=%s region=%s", *S3key, *S3secret, *S3bucket, *S3region)
}

func tearDown() {
	os.RemoveAll(drivers["disk"])
}
