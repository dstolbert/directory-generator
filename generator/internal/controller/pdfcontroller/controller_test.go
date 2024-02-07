package pdfcontroller

import (
	"testing"

	"github.com/dstolbert/osutils/osutils_mocks"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	os *osutils_mocks.Osutils
}

func buildMocks(t *testing.T) (Params, mocks) {
	params := Params{}
	return params, mocks{}
}

func TestNew(t *testing.T) {
	p, _ := buildMocks(t)
	assert.NotNil(t, p)
	New(p)
}
