// Package unstaged parses a yaml configuration file containing a list
// of repos that need to be checked for unstaged changes.
//
// The repository yaml should live in path:
//  $HOME/.unstaged.yaml
// This configuration should be in the following format.
//  repos:
//    - /path/to/repo/one
//    - /path/to/repo/two
//    - ...
//
// This script will only check for unstaged changes and currently
// does not check committed changes against an upstream.
//
// Copyright 2019 Jose R. Gonzalez. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package main

import (
	"github.com/komish/unstaged/cmd"
	"os"
)

func main() {
	os.Exit(cmd.Run())
}
