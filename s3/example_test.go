// Copyright 2014 Dynamic Design. All rights reserved.

package s3_test

import (
	"bytes"
	"io"
	"log"
	"testing"

	"github.com/dynamic-design/storage"
	_ "github.com/dynamic-design/storage/s3"
)

func TestIntegration(t *testing.T) {
	b, err := storage.Open("s3", "key=AKIAIEXPRCEXGMUEMY4A secret=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 bucket=uploadservice-test region=ap-southeast-1")
	catch(err)

	path := "sub/glados.txt"
	content := []byte("You haven't escaped, you know.\n")

	// Create the file
	f, err := b.Create(path)
	if err != nil {
		t.Error(err)
	}
	if _, err := f.Write(content); err != nil {
		t.Error(err)
	}
	if err := f.Close(); err != nil {
		t.Error(err)
	}

	// Open it
	f, err = b.Open(path)
	if err != nil {
		t.Error(err)
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, f); err != nil {
		t.Error(err)
	}
	if err := f.Close(); err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf.Bytes(), content) {
		t.Error("read file content is not the same as we stored")
	}

	// Make sure we can create an URL for the file
	if _, err := b.URL(path); err != nil {
		t.Error(err)
	}
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
