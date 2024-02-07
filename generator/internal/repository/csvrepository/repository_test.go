package csvrepository

import (
	"encoding/csv"
	"io"
	"testing"

	"github.com/dstolbert/directory-generator/entities"
	"github.com/dstolbert/osutils/osutils_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mocks struct {
	os *osutils_mocks.Osutils
}

func buildMocks(t *testing.T) (Params, mocks) {
	os := osutils_mocks.NewOsutils(t)
	params := Params{
		OS:       os,
		Filepath: "my-test.csv",
	}
	return params, mocks{os: os}
}

func TestNew(t *testing.T) {
	p, m := buildMocks(t)

	//mocks
	m.os.EXPECT().Stat(p.Filepath).Return(nil, nil)
	m.os.EXPECT().Open(p.Filepath).Return(nil, nil)
	m.os.EXPECT().CSVNewReader(mock.Anything).Return(nil)
	nCalls := 0
	m.os.EXPECT().CSVRead(mock.Anything).RunAndReturn(func(r *csv.Reader) ([]string, error) {
		// return EOF
		nCalls += 1
		if nCalls > 5 {
			return nil, io.EOF
		}

		return entities.GenerateCsvLine(), nil
	})

	assert.NotNil(t, p)
	New(p)
}
