package csvrepository

func (r *repository) Get() []Entry {
	return r.data
}
