package csvrepository

import (
	"encoding/csv"
	"io"
	"testing"

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

func generateCsvLine() []string {

	return []string{
		"LastName",
		"FirstName_Man",
		"FirstName_Woman",
		"Street",
		"City",
		"State",
		"Zip",
		"HomePhone",
		"WeddingAnniversary_Month",
		"WeddingAnniversary_Day",
		"MansEmail",
		"MansCell",
		"MansSaintName",
		"MansBirthday_Month",
		"MansBirthday_Day",
		"WomansEmail",
		"WomansCell",
		"WomansSaintName",
		"WomansBirthday_Month",
		"WomansBirthday_Day",
		"Child_1_First_Name",
		"Child_1_Saint_Name",
		"Child_1_Birthday_Month",
		"Child_1_Birthday_Day",
		"Child_2_First_Name",
		"Child_2_Saint_Name",
		"Child_2_Birthday_Month",
		"Child_2_Birthday_Day",
		"Child_3_First_Name",
		"Child_3_Saint_Name",
		"Child_3_Birthday_Month",
		"Child_3_Birthday_Day",
		"Child_4_First_Name",
		"Child_4_Saint_Name",
		"Child_4_Birthday_Month",
		"Child_4_Birthday_Day",
		"Child_5_First_Name",
		"Child_5_Saint_Name",
		"Child_5_Birthday_Month",
		"Child_5_Birthday_Day",
		"Child_6_First_Name",
		"Child_6_Saint_Name",
		"Child_6_Birthday_Month",
		"Child_6_Birthday_Day",
		"Child_7_First_Name",
		"Child_7_Saint_Name",
		"Child_7_Birthday_Month",
		"Child_7_Birthday_Day",
	}
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

		return generateCsvLine(), nil
	})

	assert.NotNil(t, p)
	New(p)
}
