package main

import (
	"github.com/eric-fleming/PDFmaker/pdfs/cars"
	"github.com/eric-fleming/PDFmaker/pdfs/fruits"
)

const (
	outputPath = "./output/"
)

func main() {

	// grab information from commandline

	// build corresponding PDF
	fruits.PrintPdf(outputPath)
	cars.PrintPdf(outputPath)
}
