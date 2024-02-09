package pdfcontroller

import (
	"errors"
	"math"
	"os"

	"github.com/dstolbert/directory-generator/entities"
	"github.com/go-pdf/fpdf"
	"github.com/sirupsen/logrus"
)

const (
	colCount = 4
	marginH  = 10.0
	lineHt   = 5.5
	photoHt  = 60.0
	cellGap  = 1.0
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
	colWeights := []float64{2, 2, 1, 3}
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
		// Format cells and determine needed height based on largest cell height
		for col := 0; col < colCount; col++ {

			cell := cellType{}

			if col == 0 && fmly.FirstName_Man != "" {
				cell.str = fmtManStr(fmly)
				photo, err = c.photos.Get(fmly.FirstName_Man, fmly.LastName)
				if err != nil && !errors.Is(err, os.ErrNotExist) {
					logrus.Errorln("error finding photo: ", err)
				}
			} else if col == 1 && fmly.FirstName_Woman != "" {
				cell.str = fmtWomanStr(fmly)

				// try and find photo if not listed by mans name
				if photo.Filepath == "" {
					photo, err = c.photos.Get(fmly.FirstName_Woman, fmly.LastName)
					if err != nil && !errors.Is(err, os.ErrNotExist) {
						logrus.Errorln("error finding photo: ", err)
					}
				}
			} else if col == 2 && fmly.Child_1_First_Name != "" {
				cell.str = fmtChildStr(fmly)
			}

			cell.wd = getColWidth(pageW, marginH, colWeights, col)
			cell.list = pdf.SplitLines([]byte(cell.str), cell.wd-cellGap-cellGap)
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[col] = cell
		}

		// Format cell with photo in it
		if photo.Filepath != "" && photoHt > maxHt {
			maxHt = photoHt
		}

		// do we need another page to fit this row?
		if y+maxHt+cellGap+cellGap+marginH+marginH > pageH {
			pdf.AddPage()
			y = pdf.GetY()
		}

		// Cell render loop
		x := marginH
		for col := 0; col < colCount; col++ {
			cell := cellList[col]
			pdf.Rect(x, y, cell.wd, maxHt+cellGap+cellGap, "D")
			cellY := y + cellGap

			// Text columns
			if col < 3 {
				for splitJ := 0; splitJ < len(cell.list); splitJ++ {
					pdf.SetXY(x+cellGap, cellY)
					pdf.CellFormat(cell.wd-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
						"L", false, 0, "")
					cellY += lineHt
				}
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
			}

			x += cell.wd

		}
		y += maxHt + cellGap + cellGap
	}

	return pdf.OutputFileAndClose(c.output)
}

func fmtManStr(family entities.Entry) string {
	str := family.FirstName_Man + " " + family.LastName + "\n"
	if family.HomePhone != "" {
		str += "home: " + family.HomePhone + "\n"
	}
	if family.MansCell != "" {
		str += "phone: " + family.MansCell + "\n"
	}
	if family.MansEmail != "" {
		str += "email: " + family.MansEmail + "\n"
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
	str := family.FirstName_Woman + " " + family.LastName + "\n"
	if family.HomePhone != "" {
		str += "home: " + family.HomePhone + "\n"
	}
	if family.WomansCell != "" {
		str += "phone: " + family.WomansCell + "\n"
	}
	if family.WomansEmail != "" {
		str += "email: " + family.WomansEmail + "\n"
	}
	if family.WomansBirthday_Month != "" {
		str += "birthday: " + family.WomansBirthday_Month + "/" + family.WomansBirthday_Day + "\n"
	}

	return str
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
