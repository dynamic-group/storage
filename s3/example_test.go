// Copyright 2014 Dynamic Design. All rights reserved.

package s3_test

import (
	"io"
	"log"
	"os"

	"github.com/dynamic-design/storage"
	_ "github.com/dynamic-design/storage/s3"
)

func ExampleFile() {
	b, err := storage.Open("s3", "key=AKIAIEXPRCEXGMUEMY4A secret=50qzh2HoqGCLFc8tcvIVRPYjl4VBYqzvSHaxxQsF bucket=uploadservice-test region=ap-southeast-1")
	catch(err)

	path := "sub/glados.txt"

	f, err := b.Create(path)
	catch(err)
	_, err = f.Write([]byte("You haven't escaped, you know.\n"))
	catch(err)
	err = f.Close()
	catch(err)

	f, err = b.Open(path)
	catch(err)
	_, err = io.Copy(os.Stdout, f)
	catch(err)
	err = f.Close()
	catch(err)

	URL, err := b.URL(path)
	catch(err)

	log.Println(URL)

	// Output: You haven't escaped, you know.
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
