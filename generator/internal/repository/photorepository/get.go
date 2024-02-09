package photorepository

import (
	"image"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/dstolbert/directory-generator/entities"
	"github.com/sirupsen/logrus"
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
