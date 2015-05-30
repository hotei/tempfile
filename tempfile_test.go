// tempfile_test.go (c) David Rook

// Package tempfile provides tools to create and delete temporary files
package tempfile

import (
	"fmt"
	"os"
	"testing"
	//
	//"github.com/hotei/tempfile"
)

func Test_001(t *testing.T) {
	defer Cleanup()
	for i := 0; i < 5; i++ {
		f, err := New("/home/mdr/tmp", "prefix", "suffix")
		if err != nil {
			fmt.Printf("tempfile create failed\n")
			os.Exit(-1)
		}
		fmt.Printf("Created %s\n", f.Name())
		DeferDelete(f.Name())
	}
}
