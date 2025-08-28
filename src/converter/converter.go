package converter

import (
    "fmt"
    "os/exec"
    "path/filepath"
    "svg-to-png-converter/src/types"
)

// ConvertSVGToPNG converts an SVG file to a PNG file using ImageMagick's convert command.
func ConvertSVGToPNG(svgPath string, opts *types.Options) error {
    pngPath := filepath.Dir(svgPath) + "/" + basename(svgPath) + ".png"

    cmd := exec.Command("inkscape", "--export-filename="+pngPath, svgPath)

	if opts.Debug {
        fmt.Println("Running:", cmd.String())
    }
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("ImageMagick convert failed: %w", err)
    }
    return nil
}

func basename(filePath string) string {
    fileName := filepath.Base(filePath)
    extension := filepath.Ext(fileName)
    return fileName[0 : len(fileName)-len(extension)]
}