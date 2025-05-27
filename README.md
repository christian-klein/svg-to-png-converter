# SVG to PNG Converter

A simple Go application that recursively searches a directory (and its subdirectories) for SVG files and converts them to PNG format.

## Features

- Recursively walks through a given directory
- Converts all `.svg` files to `.png` using [oksvg](https://github.com/srwiley/oksvg) and [rasterx](https://github.com/srwiley/rasterx)
- Optional debug output for verbose logging

## Requirements

- Go 1.19 or newer

## Installation

Clone the repository and install dependencies:

```bash
git clone https://github.com/christian-klein/svg-to-png-converter.git
cd svg-to-png-converter
go mod tidy
```

## Usage

Run the application with:

```bash
go run src/main.go --path /path/to/svg/files --debug
```

### Command-line Flags

- `--path` or `-p`: Root path to start searching for SVG files (default: current directory)
- `--debug` or `-d`: Enable debug output

**Example:**

```bash
 go run src/main.go -p ./sample -d       
```

## Output

- Each SVG file found will be converted to a PNG file in the same directory as the original SVG.

## Project Structure

```
svg-to-png-converter/
├── src/
│   ├── main.go
│   ├── types/
│   │   └── types.go
│   ├── utils/
│   │   └── filewalker.go
│   └── converter/
│       └── converter.go
├── go.mod
├── go.sum
└── README.md
```

## License

MIT License

---

**Author:** Christian Klein