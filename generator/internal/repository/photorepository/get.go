package photorepository

import (
	"os"
	"strings"
)

func (r *repository) Get(firstName, lastName string) (string, error) {

	// search photo dir for a filename match
	files, err := r.os.ReadDir("./")
	if err != nil {
		return "", err
	}

	for _, f := range files {

		// substring match last and first names
		if strings.Contains(f.Name(), firstName) && strings.Contains(f.Name(), lastName) {
			return f.Name(), nil
		}

	}

	return "", os.ErrNotExist
}
