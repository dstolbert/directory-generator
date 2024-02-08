package photorepository

import (
	"testing"

	"github.com/dstolbert/osutils/osutils_mocks"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	os *osutils_mocks.Osutils
}

func buildMocks(t *testing.T) (Params, mocks) {
	os := osutils_mocks.NewOsutils(t)
	params := Params{
		OS:       os,
		Filepath: "data/photos",
	}
	return params, mocks{os: os}
}

func TestNew(t *testing.T) {
	p, m := buildMocks(t)

	//mocks
	m.os.EXPECT().Stat(p.Filepath).Return(nil, nil)
	assert.NotNil(t, p)
	New(p)
}
