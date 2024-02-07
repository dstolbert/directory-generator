package csvrepository

import (
	"errors"
	"io"
	"os"

	"github.com/dstolbert/osutils"
	"github.com/sirupsen/logrus"
)

//go:generate mockery --config $WD/test/.mockery.yml --all --output $PWD/mocks --outpkg csvrepository_mock
type Repository interface {
}

type repository struct {
	filepath string
	os       osutils.Osutils
	data     []Entry
}

type Params struct {
	Filepath string
	OS       osutils.Osutils
}

func (r *repository) init() {

	// does the file exist?
	if _, err := r.os.Stat(r.filepath); errors.Is(err, os.ErrNotExist) {
		logrus.Panicf("could not initialize repo: %s", err.Error())
	}

	// read file
	f, err := r.os.Open(r.filepath)
	if err != nil {
		logrus.Panicf("could not read repo data: %s", err.Error())
	}
	defer f.Close()

	// Load csv
	r.data = []Entry{}
	csvReader := r.os.CSVNewReader(f)
	i := 0
	for {
		rec, err := r.os.CSVRead(csvReader)
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Panicf("error parsing csv: %s", err.Error())
		}

		// parse data to Entry struct
		if i > 0 {
			r.data = append(r.data, parseLine(rec))
		}
		i++
	}
}

func New(p Params) Repository {
	r := repository{
		filepath: p.Filepath,
		os:       p.OS,
	}
	r.init()

	return &r
}
