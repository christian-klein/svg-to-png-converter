package converter

import (
	"fmt"
	"image"

	"image/png"
	"os"
	"path/filepath"
	"svg-to-png-converter/src/types"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

// ConvertSVGToPNG converts an SVG file to a PNG file.
func ConvertSVGToPNG(svgPath string, opts *types.Options) error {
	svg, _ := oksvg.ReadIcon(svgPath, oksvg.WarnErrorMode)
	width := int(svg.ViewBox.W)
	height := int(svg.ViewBox.H)

	if opts.Debug {
		fmt.Println("\nProcessing SVG file:", svgPath)
		fmt.Println("SVG ViewBox Width:", width, "Height:", height)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	scanner := rasterx.NewScannerGV(width, height, img, img.Bounds())
	raster := rasterx.NewDasher(width, height, scanner)

	svg.SetTarget(0, 0, float64(width), float64(height))

	opacity := 1.0
	svg.Draw(raster, opacity)
	
	name := filepath.Dir(svgPath) + "/" + basename(svgPath) + ".png"

	if opts.Debug {
		fmt.Println("Saving PNG file:", name)
	}

	// Save the PNG file
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	return nil
}

func basename(filePath string) string {
	fileName := filepath.Base(filePath)
	extension := filepath.Ext(fileName)
	return fileName[0 : len(fileName)-len(extension)]
}