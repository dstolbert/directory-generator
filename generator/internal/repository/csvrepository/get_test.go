package csvrepository

import (
	"encoding/csv"
	"fmt"
	"io"
	"testing"

	"github.com/dstolbert/directory-generator/entities"
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
			data := r.Get()
			fmt.Println(data)
		})
	}
}
