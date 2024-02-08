package photorepository

import (
	"os"
	"path/filepath"
	"strings"
)

func (r *repository) Get(firstName, lastName string) (string, error) {

	// lowercase names
	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)

	// search photo dir for a filename match
	files, err := r.os.ReadDir(r.filepath)
	if err != nil {
		return "", err
	}

	for _, f := range files {

		// substring match last and first names
		lower := strings.ToLower(f.Name())
		if strings.Contains(lower, firstName) && strings.Contains(lower, lastName) {
			return filepath.Join(r.filepath, f.Name()), nil
		}

	}

	return "", os.ErrNotExist
}
