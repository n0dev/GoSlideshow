package exif

import (
	"fmt"
	"os"
)

// Read displays exif information about the file
func Read(f *os.File) {
	fmt.Println("Exif information")
}