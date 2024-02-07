package pdfcontroller

import (
	"fmt"

	"github.com/dstolbert/directory-generator/entities"
	"github.com/jung-kurt/gofpdf"
)

const (
	colCount = 3
	colWd    = 60.0
	marginH  = 10.0
	lineHt   = 5.5
	cellGap  = 1.0
)

// var colStrList [colCount]string
type cellType struct {
	str  string
	list [][]byte
	ht   float64
}

func (c *controller) SavePDF() error {

	// get data from repo
	data := c.csv.Get()

	pdf := gofpdf.New("P", "mm", "A4", "") // 210 x 297
	alignList := [colCount]string{"L", "L", "L"}
	pdf.SetFont("Arial", "", 10)
	pdf.AddPage()

	_, pageH := pdf.GetPageSize()

	// Rows
	y := pdf.GetY()
	for _, fmly := range data {

		maxHt := lineHt
		cellList := [colCount]cellType{}
		// Cell height calculation loop
		for col := 0; col < colCount; col++ {

			cell := cellType{}

			if col == 0 && fmly.FirstName_Man != "" {
				cell.str = fmtManStr(fmly)
			} else if col == 1 && fmly.FirstName_Woman != "" {
				cell.str = fmtWomanStr(fmly)
			} else if col == 2 && fmly.Child_1_First_Name != "" {
				cell.str = fmtChildStr(fmly)
			}

			cell.list = pdf.SplitLines([]byte(cell.str), colWd-cellGap-cellGap)
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[col] = cell
		}

		// do we need another page to fit this row?
		if y+maxHt+cellGap+cellGap+marginH+marginH > pageH {
			pdf.AddPage()
			y = pdf.GetY()
			fmt.Println("adding page. new y: ", y)
		}

		// Cell render loop
		x := marginH
		for col := 0; col < colCount; col++ {
			pdf.Rect(x, y, colWd, maxHt+cellGap+cellGap, "D")
			cell := cellList[col]
			cellY := y + cellGap
			for splitJ := 0; splitJ < len(cell.list); splitJ++ {
				pdf.SetXY(x+cellGap, cellY)
				pdf.CellFormat(colWd-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
					alignList[col], false, 0, "")
				cellY += lineHt
			}
			x += colWd
		}
		y += maxHt + cellGap + cellGap
	}

	return pdf.OutputFileAndClose(c.output)
}

func fmtManStr(family entities.Entry) string {
	str := family.FirstName_Man + " " + family.LastName + "\n" +
		"home: " + family.HomePhone + "\n" +
		"phone: " + family.MansCell + "\n" +
		"email: " + family.MansEmail + "\n" +
		"birthday: " + family.MansBirthday_Month + "/" + family.MansBirthday_Day + "\n"

	if family.WeddingAnniversary_Month != "" {
		str += "anniversary:" + family.WeddingAnniversary_Month + "/" + family.WeddingAnniversary_Day
	}

	return str
}

func fmtWomanStr(family entities.Entry) string {
	return family.FirstName_Woman + " " + family.LastName + "\n" +
		"home: " + family.HomePhone + "\n" +
		"phone: " + family.WomansCell + "\n" +
		"email: " + family.WomansEmail + "\n" +
		"birthday: " + family.WomansBirthday_Month + "/" + family.WomansBirthday_Day
}

func fmtChildStr(family entities.Entry) string {
	str := ""

	if family.Child_1_First_Name != "" {
		str += family.Child_1_First_Name + " " + family.Child_1_Birthday_Month + "/" + family.Child_1_Birthday_Day + "\n"
	}
	if family.Child_2_First_Name != "" {
		str += family.Child_2_First_Name + " " + family.Child_2_Birthday_Month + "/" + family.Child_2_Birthday_Day + "\n"
	}
	if family.Child_3_First_Name != "" {
		str += family.Child_3_First_Name + " " + family.Child_3_Birthday_Month + "/" + family.Child_3_Birthday_Day + "\n"
	}
	if family.Child_4_First_Name != "" {
		str += family.Child_4_First_Name + " " + family.Child_4_Birthday_Month + "/" + family.Child_4_Birthday_Day + "\n"
	}
	if family.Child_5_First_Name != "" {
		str += family.Child_5_First_Name + " " + family.Child_5_Birthday_Month + "/" + family.Child_5_Birthday_Day + "\n"
	}
	if family.Child_6_First_Name != "" {
		str += family.Child_6_First_Name + " " + family.Child_6_Birthday_Month + "/" + family.Child_6_Birthday_Day + "\n"
	}
	if family.Child_7_First_Name != "" {
		str += family.Child_7_First_Name + " " + family.Child_7_Birthday_Month + "/" + family.Child_7_Birthday_Day
	}

	return str
}