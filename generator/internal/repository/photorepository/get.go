package photorepository

import (
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/dstolbert/directory-generator/entities"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/draw"
)

const (
	maxDim = 600
)

func (r *repository) Get(firstName, lastName string) (entities.Image, error) {

	// format names for filename matching
	firstName = cleanName(firstName)
	lastName = cleanName(lastName)
	img := entities.Image{}

	// search photo dir for a filename match
	files, err := r.os.ReadDir(r.filepath)
	if err != nil {
		return img, err
	}

	for _, f := range files {

		// substring match last and first names
		lower := strings.ToLower(f.Name())
		if strings.Contains(lower, firstName) && strings.Contains(lower, lastName) {
			img.Filepath = filepath.Join(r.filepath, f.Name())

			// load image to determine aspect ratio
			if reader, err := os.Open(img.Filepath); err == nil {
				defer reader.Close()
				im, _, err := image.DecodeConfig(reader)
				if err != nil {
					logrus.Errorf("%s: %v\n", img.Filepath, err)
					continue
				}
				img.Height = im.Height
				img.Width = im.Width

				// resize image if needed
				if img.Height > maxDim || img.Width > maxDim {
					input, _ := os.Open(img.Filepath)
					defer input.Close()

					// Decode the image (from jpg to image.Image):
					src, _ := jpeg.Decode(input)

					output, _ := os.Create(img.Filepath)
					defer output.Close()

					// Set the expected size that you want:
					dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))

					// Resize:
					draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

					// Encode to `output`:
					jpeg.Encode(output, dst, &jpeg.Options{Quality: 90})
				}
			} else {
				logrus.Errorln("Impossible to open the file:", err)
			}

			return img, nil
		}

	}

	return img, os.ErrNotExist
}

func cleanName(name string) string {
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)

	// Now remove all non-alpha chars
	chars := " abcdefghijklmnopqrstuvwxyz"
	filtered := []rune{}
	for _, r := range name {
		if strings.ContainsRune(chars, r) {
			filtered = append(filtered, r)
		}
	}
	return string(filtered)
}
