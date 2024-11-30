package ui

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
	"strings"
)

func (paths *PathsToResumes) GetHtmlToPDF(htmlFilePath, outputPath string) error {
	info, err := os.Stat(paths.GeneratedResumePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Последняя модификация файла:", info.ModTime())
	htmlFile, err := os.ReadFile(paths.GeneratedResumePath)
	if err != nil {
		return err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(string(htmlFile)))

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(paths.ConvertedToPdfPath)
	log.Printf("PDF успешно сохранен в %s\n", paths.ConvertedToPdfPath)
	return err
}
