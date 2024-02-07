package pdfcontroller

import "fmt"

func (c *controller) SavePDF() error {

	// get data from repo
	data := c.csv.Get()

	// save to file
	fmt.Println(data)

	return nil
}
