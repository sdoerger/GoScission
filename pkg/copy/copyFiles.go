package copy

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"

	make "github.com/sdoerger/GoScission/pkg/make"
)

func CopyFiles(s int /* start index */, e int /* end index */, d string /* dir path */, f []fs.FileInfo /* files */, i int /* index for dir */, targetDir string) {
	// s = START
	// e = END

	fmt.Println(targetDir)

	dn := "dir_" + strconv.Itoa(i+1) // dir name + index

	// CREATE: Dir for current file range
	make.MakeDir(d+"/"+targetDir, dn)

	// Current File Selection
	cfs := f[s:e]

	for _, file := range cfs {
		// Exclude destination dir
		f := file.Name()

		if f != targetDir {
			// fmt.Println(f)

			// Open original file
			original, err := os.Open(d + "/" + f)
			if err != nil {
				log.Fatal(err)
			}

			defer original.Close()

			// Get file name
			fn := filepath.Base(original.Name())

			// Create new file
			new, err := os.Create(d + "/" + targetDir + "/" + dn + "/" + fn)
			if err != nil {
				log.Fatal(err)
			}
			defer new.Close()

			//This will copy
			bytesWritten, err := io.Copy(new, original)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Bytes Written: %d\n", bytesWritten)

		}

	}

}
