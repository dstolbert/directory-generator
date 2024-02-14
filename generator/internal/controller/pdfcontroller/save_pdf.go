package pdfcontroller

import (
	"errors"
	"fmt"
	"math"
	"os"

	"github.com/dstolbert/directory-generator/entities"
	"github.com/go-pdf/fpdf"
	"github.com/sirupsen/logrus"
)

const (
	colCount       = 4
	marginH        = 4.0
	lineHt         = 5.5
	photoHt        = 50.0
	cellGap        = 1.0
	familyHeaderHt = 10.0
)

// var colStrList [colCount]string
type cellType struct {
	str  string
	list [][]byte
	ht   float64
	wd   float64
}

func (c *controller) SavePDF() error {

	// get data from repo
	data := c.csv.Get()
	pdf := fpdf.New("P", "mm", "A4", "") // 210 x 297
	colWeights := []float64{2, 2, 2, 1}
	pdf.SetFont("Arial", "", 10)
	pdf.AddPage()

	pageW, pageH := pdf.GetPageSize()

	// Rows
	y := pdf.GetY()
	for _, fmly := range data {

		maxHt := lineHt
		cellList := [colCount]cellType{}
		photo := entities.Image{}
		var err error
		totalColWid := 0.0
		// Format cells and determine needed height based on largest cell height
		for col := 0; col < colCount; col++ {

			cell := cellType{}

			if col == 1 && fmly.FirstName_Man != "" {
				cell.str = fmtManStr(fmly)
				photo, err = c.photos.Get(fmly.FirstName_Man, fmly.LastName)
				if err != nil && !errors.Is(err, os.ErrNotExist) {
					logrus.Errorln("error finding photo: ", err)
				}
			} else if col == 2 && fmly.FirstName_Woman != "" {
				cell.str = fmtWomanStr(fmly)

				// try and find photo if not listed by mans name
				if photo.Filepath == "" {
					photo, err = c.photos.Get(fmly.FirstName_Woman, fmly.LastName)
					if err != nil && !errors.Is(err, os.ErrNotExist) {
						logrus.Errorln("error finding photo: ", err)
					}
				}
			} else if col == 3 && len(fmly.Children) > 0 {
				cell.str = fmtChildStr(fmly)
			}

			cell.wd = getColWidth(pageW, marginH, colWeights, col)
			totalColWid += cell.wd
			cell.list = pdf.SplitLines([]byte(cell.str), cell.wd-cellGap-cellGap)
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[col] = cell
		}

		// if no mans name, shift cells, recalculate cell widths
		if fmly.FirstName_Man == "" {
			cellList[0] = cellList[1]
			cellList[0].wd = getColWidth(pageW, marginH, colWeights[1:], 0)
			cellList[1] = cellList[2]
			cellList[1].wd = getColWidth(pageW, marginH, colWeights[1:], 1)
			cellList[2] = cellList[3]
			cellList[2].wd = getColWidth(pageW, marginH, colWeights[1:], 2)
			cellList[3] = cellType{}
		} else if fmly.FirstName_Woman == "" {
			cellList[1].wd = getColWidth(pageW, marginH, colWeights[1:], 0)
			cellList[2] = cellList[3]
			cellList[2].wd = getColWidth(pageW, marginH, colWeights[1:], 1)
			cellList[3] = cellType{}
		}

		// Format cell with photo in it
		if photo.Filepath != "" && photoHt > maxHt {
			maxHt = photoHt
		}

		// do we need another page to fit this row?
		if y+maxHt+cellGap+cellGap+marginH+marginH+familyHeaderHt+10 > pageH {
			pdf.AddPage()
			y = pdf.GetY()
		}

		// Create family name header
		x := marginH
		pdf.SetFillColor(194, 196, 195)
		pdf.SetFont("Arial", "B", 16)
		pdf.Rect(x, y, totalColWid, familyHeaderHt, "DF")
		cellY := y + cellGap
		pdf.SetXY(x+cellGap, cellY)
		pdf.CellFormat(totalColWid, familyHeaderHt, fmly.LastName, "", 0,
			"C", false, 0, "")
		y += familyHeaderHt + 1
		pdf.SetFont("Arial", "", 10)

		// Cell render loop
		pdf.SetFillColor(256, 256, 256)
		for col := 0; col < colCount; col++ {
			cell := cellList[col]
			pdf.Rect(x, y, cell.wd, maxHt+cellGap+cellGap, "F")
			cellY := y + cellGap

			// Text columns
			if col > 0 {
				for splitJ := 0; splitJ < len(cell.list); splitJ++ {
					pdf.SetXY(x+cellGap, cellY)
					pdf.CellFormat(cell.wd-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
						"L", false, 0, "")
					cellY += lineHt
				}
				x += cell.wd
			} else if photo.Filepath != "" {

				// normally, we set height and dynamically scale width
				ht := photoHt
				wt := 0.0

				// if image is horizontal, then scale width and do the height dynamically
				if photo.Width > photo.Height {
					ht = 0
					wt = cell.wd - cellGap - cellGap
				}

				// Photo columns
				pdf.ImageOptions(photo.Filepath, x+cellGap, cellY, wt, ht, false, fpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				}, 0, "S")
				pdf.GetImageInfo(photo.Filepath).SetDpi(1)
				x += cell.wd
			}

		}
		y += maxHt + cellGap + cellGap
	}

	return pdf.OutputFileAndClose(c.output)
}

func fmtManStr(family entities.Entry) string {
	str := family.FirstName_Man + "\n"
	if family.HomePhone != "" {
		str += family.HomePhone + "\n"
	}
	if family.MansCell != "" && family.MansCell != family.HomePhone {
		str += "cell: " + family.MansCell + "\n"
	}
	if family.MansEmail != "" {
		str += family.MansEmail + "\n"
	}
	if family.MansBirthday_Month != "" {
		str += "birthday: " + family.MansBirthday_Month + "/" + family.MansBirthday_Day + "\n"
	}
	if family.WeddingAnniversary_Month != "" {
		str += "anniversary:" + family.WeddingAnniversary_Month + "/" + family.WeddingAnniversary_Day
	}

	return str
}

func fmtWomanStr(family entities.Entry) string {
	str := family.FirstName_Woman + "\n"
	if family.HomePhone != "" && family.FirstName_Man == "" {
		str += family.HomePhone + "\n"
	}
	if family.WomansCell != "" && family.WomansCell != family.HomePhone {
		str += "cell: " + family.WomansCell + "\n"
	}
	if family.WomansEmail != "" {
		str += family.WomansEmail + "\n"
	}
	if family.WomansBirthday_Month != "" {
		str += "birthday: " + family.WomansBirthday_Month + "/" + family.WomansBirthday_Day + "\n"
	}

	return str
}

func fmtChildStr(family entities.Entry) string {

	str := ""

	for _, child := range family.Children {
		fmt.Println("adding child: ", child.FirstName, family.LastName)
		str += child.FirstName
		if child.BirthdayMonth != "" {
			str += " " + child.BirthdayMonth + "/" + child.BirthdayDay
		}
		str += "\n"
	}

	return str
}

func getColWidth(pageW, margin float64, colWeights []float64, colIndex int) float64 {
	// subtract margins
	pageW -= (margin * 2)

	// scale col based on weight
	totalWeight := 0.0
	for _, w := range colWeights {
		totalWeight += w
	}
	return math.Floor(pageW * (colWeights[colIndex] / totalWeight))
}
