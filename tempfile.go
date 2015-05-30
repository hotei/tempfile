// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Not optimal for New to start with 1 every time. We decouple loop and file index.
// Changes (c) 2015 David Rook all rights reserved. Released with BSD 2-clause
// license.

// Package tempfile provides tools to create and delete temporary files
package tempfile

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const maxIndex = 10000

var (
	lastIndex int
	tempFiles []string
	tempFilesMu = sync.Mutex{}
)

// New returns an unused filename for output files.
func New(dir, prefix, suffix string) (*os.File, error) {
	loopCt:=0
	index := lastIndex+1	// first file created will be "1" like original
	for {
		if loopCt >= maxIndex {	// enforces upper limit on directory size
			break
		}  
		if index >= maxIndex { // reuse lower numbers but only if they've been deleted
			index = 1
		}
		path := filepath.Join(dir, fmt.Sprintf("%s%03d%s", prefix, index, suffix))
		if _, err := os.Stat(path); err != nil {
			// TODO(mdr):
			// should probably test for the right error here and stop if isn't a "not exists" err
			lastIndex = index		// saves value of last file created
			return os.Create(path)
		}
		loopCt++
		index++
	}
	// Give up
	return nil, fmt.Errorf("could not create file of the form %s%03d%s", prefix, 1, suffix)
}


// DeferDelete marks a file to be deleted by next call to Cleanup()
func DeferDelete(path string) {
	tempFilesMu.Lock()
	tempFiles = append(tempFiles, path)
	tempFilesMu.Unlock()
}

// Cleanup removes any temporary files selected for deferred cleaning.
func Cleanup() {
	tempFilesMu.Lock()
	for _, f := range tempFiles {
		os.Remove(f)
	}
	tempFiles = nil
	tempFilesMu.Unlock()
}
