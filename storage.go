// Copyright 2014 Dynamic Design. All rights reserved.

// Package storage provides a single interface for various cloud storage services through the use of drivers that wrap their well known client libraries.
package storage

import (
	"fmt"

	"github.com/dynamic-design/storage/driver"
)

var drivers = make(map[string]driver.Driver)

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil, it panics.
func Register(name string, driver driver.Driver) {
	if driver == nil {
		panic("storage: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("storage: Register called twice for driver " + name)
	}

	drivers[name] = driver
}

// Open opens a bucket specified by its driver name and a driver-specific source string.
func Open(name, source string) (driver.Bucket, error) {
	driver, ok := drivers[name]
	if !ok {
		return nil, fmt.Errorf("storage: unknown driver %q (forgotten import?)", name)
	}

	return driver.Open(source)
}
