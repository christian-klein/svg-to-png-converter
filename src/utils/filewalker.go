package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"svg-to-png-converter/src/converter"
	"svg-to-png-converter/src/types"
)

// WalkDirectory recursively walks through the specified directory and its subdirectories,
// converting all SVG files to PNG format.
func WalkDirectory(opts *types.Options) error {
	filecount := 0
	err := filepath.Walk(opts.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}


		// Check if the file is an SVG file
		if !info.IsDir() && filepath.Ext(path) == ".svg" {
			filecount++
			// Convert the SVG file to PNG
			if err := converter.ConvertSVGToPNG(path, opts); err != nil {
				return err
			}
		}
		return nil
	})
	fmt.Println("Processed ", filecount, " SVG files in directory:", opts.RootDir)
	return err
}