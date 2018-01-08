// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Mon,  8 Jan 2018]

package main

import (
    "fmt"
	"github.com/eddiefisher/pinger/src/version"
)

func main() {
    fmt.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
}
