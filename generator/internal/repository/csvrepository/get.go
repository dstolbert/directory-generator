package csvrepository

import "github.com/dstolbert/directory-generator/entities"

func (r *repository) Get() []entities.Entry {
	return r.data
}
