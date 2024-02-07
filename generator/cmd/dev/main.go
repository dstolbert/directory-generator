package main

import (
	"os"

	"github.com/dstolbert/directory-generator/internal/controller/pdfcontroller"
	"github.com/dstolbert/directory-generator/internal/repository/csvrepository"
	"github.com/dstolbert/osutils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// init csv repo
	csv := csvrepository.New(csvrepository.Params{
		Filepath: os.Getenv("CSV_PATH"),
		OS:       osutils.New(osutils.Params{}),
	})

	// create controller and save data to pdf
	c := pdfcontroller.New(pdfcontroller.Params{
		CSV:    csv,
		Output: os.Getenv("PDF_OUT"),
	})
	c.SavePDF()
}
