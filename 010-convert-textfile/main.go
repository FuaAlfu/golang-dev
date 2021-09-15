package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	file := "test.txt"

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("%s file not found", file)
	}

	pdf := gofpdf.New("p", "mm", "A4", "") //p -> orientationStr :: mm -> milimeter
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	pdf.MultiCell(190, 6, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose("test.pdf")

	fmt.Println("PDF Created")
}
