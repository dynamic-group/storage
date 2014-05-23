// Copyright 2014 Dynamic Design. All rights reserved.

package disk_test

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/dynamic-design/storage"
	_ "github.com/dynamic-design/storage/disk"
)

func ExampleFile() {
	dir, err := ioutil.TempDir("", "")
	catch(err)
	defer os.RemoveAll(dir)

	b, err := storage.Open("disk", dir)
	catch(err)

	f, err := b.Create("glados.txt")
	catch(err)
	_, err = f.Write([]byte("You haven't escaped, you know.\n"))
	catch(err)
	err = f.Close()
	catch(err)

	f, err = b.Open("glados.txt")
	catch(err)
	_, err = io.Copy(os.Stdout, f)
	catch(err)
	err = f.Close()
	catch(err)

	// Output: You haven't escaped, you know.
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
