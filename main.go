package main

import (
	"log"
	"os"

	"github.com/eric-fleming/PDFmaker/pdfs/cars"
	"github.com/eric-fleming/PDFmaker/pdfs/fruits"
)

const (
	outputPath = "./output/"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalln("You have too few arguements\n1st arg [pdftype]\n2nd arg [filename]")
		os.Exit(1)
	} else if len(os.Args) > 3 {
		log.Fatalln("You have too many arguements\n1st arg [pdftype]\n2nd arg [filename]")
		os.Exit(1)
	}

	// grab information from command line
	log.Println("PDF Maker Running!")
	args := os.Args[1:]
	pdfType := args[0]
	filename := args[1]

	if filename == "" {
		log.Fatalln("Specify a file name")
		os.Exit(2)
	}

	switch pdfType {
	case "fruits":
		fruits.PrintPdf(outputPath, filename+".pdf")
	case "cars":
		cars.PrintPdf(outputPath, filename+".pdf")
	default:
		log.Println("Do not recognize this type")
		os.Exit(1)
	}

	log.Println("Complete")
}
