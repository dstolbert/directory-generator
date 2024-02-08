package photorepository

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	tcs := []struct {
		name             string
		givenAndExpected func(p Params, m *mocks)
	}{
		{
			"test get data",
			func(p Params, m *mocks) {
				//mocks
				m.os.EXPECT().Stat(p.Filepath).Return(nil, nil)
				m.os.EXPECT().ReadDir(mock.Anything).Return([]fs.DirEntry{}, nil)
			},
		},
	}

	// Run tests
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			p, m := buildMocks(t)
			tc.givenAndExpected(p, &m)
			r := New(p)
			r.Get("john", "doe")
		})
	}
}
