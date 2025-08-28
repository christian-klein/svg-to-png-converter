package main

import (
	"fmt"
	"log"
	"svg-to-png-converter/src/types"
	"svg-to-png-converter/src/utils"
	"time"

	"github.com/spf13/pflag"
)

func main() {

	// rootDir := "./" // Change this to the desired directory
	start := time.Now()
	path := pflag.StringP("path", "p", "", "Root path to start the search for SVG files")
	debug := pflag.BoolP("debug", "d", false, "Debug")
	pflag.Parse()

	opts := &types.Options{
		Debug: false,
		RootDir: "./",
	}

	if *path != "" {
		opts.RootDir = *path
	} else {
		if *debug {
			fmt.Println("No path provided, using current directory as root.")
		}
	}

	if *debug {
		fmt.Println("Debug mode is enabled.")
		opts.Debug = true
	}

	// Walk through the directory and convert SVG files to PNG
	err := utils.WalkDirectory(opts)
	if err != nil {
		log.Fatalf("Error walking the directory: %v", err)
	}

	fmt.Println("\n---------------------------")
	
	t := time.Now()
	elapsed := t.Sub(start)

	if *debug {
		fmt.Println("Elapsed:", elapsed)
	}

	fmt.Println("SVG to PNG conversion completed.")
}
