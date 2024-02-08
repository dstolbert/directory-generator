package main

import (
	"github.com/dstolbert/directory-generator/internal/controller/pdfcontroller"
	"github.com/dstolbert/directory-generator/internal/repository/csvrepository"
	"github.com/dstolbert/directory-generator/internal/repository/photorepository"
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
	os := osutils.New(osutils.Params{})
	csv := csvrepository.New(csvrepository.Params{
		Filepath: os.Getenv("CSV_PATH"),
		OS:       os,
	})

	// init photos repo
	photos := photorepository.New(photorepository.Params{
		Filepath: os.Getenv("PHOTO_DIR"),
		OS:       os,
	})

	// create controller and save data to pdf
	c := pdfcontroller.New(pdfcontroller.Params{
		CSV:    csv,
		Output: os.Getenv("PDF_OUT"),
		Photos: photos,
	})
	c.SavePDF()
}
