// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package consolidate is used to consolidate constants, types and functions
// from the sys generated files grouping them by GOOS and GOARCH.
// The files named as <name>_<goos>_<goarch>.go are processed and the consolidated
// objects are stored in a new file named <name>_<goos>.go.
//
// The Registry is used to collect all the files data, agggregate the objects and
// output the new files.
// It works as follow:
//  - parse the sys package only selecting the files we are interested in
//  - group the files by name and for each name by GOOS and GOARCH
//    (cf. the Registry type)
//  - build the list of constants, types and functions per file (GOOS/GOARCH)
//  - consolidate the objects: they must all be present in the related GOOS/GOARCH files
//    (intersection of all the constants, types and functions)
//  - remove the consolidated objects from each GOOS/GOARCG file
//  - update the input files with their stripped down version
//  - create the new files containing the consolidated objects
package consolidate
