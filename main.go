package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"os"

	_ "golang.org/x/image/webp"
)

type Color string

const (
	ColorRed   = "\u001b[31m"
	ColorGreen = "\u001b[32m"
	ColorReset = "\u001b[0m"
)

func alert(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func convertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)

	if err != nil {
		return err
	}

	return png.Encode(w, img)
}

func main() {
	sourceFile, outputFile := "", ""

	flag.StringVar(&sourceFile, "src", "", "enter source file name")
	flag.StringVar(&outputFile, "out", "", "enter output file name")

	flag.Parse()

	fmt.Println(sourceFile, outputFile)

	if sourceFile == "" || outputFile == "" {
		alert(ColorRed, "sorry, please check that you have entered all files")
		return
	}

	fmt.Printf("Converting %v to %v", sourceFile, outputFile)

	srcFile, err := os.Open(sourceFile)

	if err != nil {
		panic(err)
	}

	defer srcFile.Close()

	outFile, err := os.Create(outputFile)

	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	err = convertToPNG(outFile, srcFile)

	if err != nil {
		panic(err)
	}

	alert(ColorGreen, "Image sucessfully converted to png!")
}
