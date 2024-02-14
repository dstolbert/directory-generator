package entities

type Entry struct {
	LastName                 string
	FirstName_Man            string
	FirstName_Woman          string
	Street                   string
	City                     string
	State                    string
	Zip                      string
	HomePhone                string
	WeddingAnniversary_Month string
	WeddingAnniversary_Day   string
	MansEmail                string
	MansCell                 string
	MansSaintName            string
	MansBirthday_Month       string
	MansBirthday_Day         string
	WomansEmail              string
	WomansCell               string
	WomansSaintName          string
	WomansBirthday_Month     string
	WomansBirthday_Day       string
	Children                 []Child
}

type Child struct {
	FirstName     string
	SaintName     string
	BirthdayMonth string
	BirthdayDay   string
}

// helper to convert a line from a csv to an entry
func ParseLine(line []string) Entry {
	e := Entry{}

	// big nasty switch
	for i, v := range line {
		switch i {
		case 0:
			e.LastName = v
		case 1:
			e.FirstName_Man = v
		case 2:
			e.FirstName_Woman = v
		case 3:
			e.Street = v
		case 4:
			e.City = v
		case 5:
			e.State = v
		case 6:
			e.Zip = v
		case 7:
			e.HomePhone = v
		case 8:
			e.WeddingAnniversary_Month = v
		case 9:
			e.WeddingAnniversary_Day = v
		case 10:
			e.MansEmail = v
		case 11:
			e.MansCell = v
		case 12:
			e.MansSaintName = v
		case 13:
			e.MansBirthday_Month = v
		case 14:
			e.MansBirthday_Day = v
		case 15:
			e.WomansEmail = v
		case 16:
			e.WomansCell = v
		case 17:
			e.WomansSaintName = v
		case 18:
			e.WomansBirthday_Month = v
		case 19:
			e.WomansBirthday_Day = v
		case 20:
			saveChild(Child{
				FirstName:     v,
				SaintName:     line[i+1],
				BirthdayMonth: line[i+2],
				BirthdayDay:   line[i+3],
			}, &e)
		case 24:
			saveChild(Child{
				FirstName:     v,
				SaintName:     line[i+1],
				BirthdayMonth: line[i+2],
				BirthdayDay:   line[i+3],
			}, &e)

		case 28:
			saveChild(Child{
				FirstName:     v,
				SaintName:     line[i+1],
				BirthdayMonth: line[i+2],
				BirthdayDay:   line[i+3],
			}, &e)

		case 32:
			saveChild(Child{
				FirstName:     v,
				SaintName:     line[i+1],
				BirthdayMonth: line[i+2],
				BirthdayDay:   line[i+3],
			}, &e)

		case 36:
			saveChild(Child{
				FirstName:     v,
				SaintName:     line[i+1],
				BirthdayMonth: line[i+2],
				BirthdayDay:   line[i+3],
			}, &e)

		case 40:
			saveChild(Child{
				FirstName:     v,
				SaintName:     line[i+1],
				BirthdayMonth: line[i+2],
				BirthdayDay:   line[i+3],
			}, &e)
		case 44:
			saveChild(Child{
				FirstName:     v,
				SaintName:     line[i+1],
				BirthdayMonth: line[i+2],
				BirthdayDay:   line[i+3],
			}, &e)

		}
	}

	return e
}

// Saves child in place to the family.Children array
func saveChild(c Child, family *Entry) {

	if family.Children == nil {
		family.Children = []Child{}
	}

	// Only save if child has a name
	if c.FirstName != "" {
		family.Children = append(family.Children, c)
	}
}

// Simple helper to generate test data
func GenerateCsvLine() []string {

	return []string{
		"Smith",
		"John",
		"Jane",
		"123 Fun ave",
		"Mesa",
		"AZ",
		"85123",
		"480-123-4567",
		"February",
		"7",
		"john.smith@fun.com",
		"480-123-4567",
		"St John",
		"Jan",
		"1",
		"jane.smith@fun.com",
		"602-123-4567",
		"St Mary",
		"Dec",
		"4th",
		"Billy",
		"St John",
		"Mar",
		"10th",
		"Willy",
		"St Paul",
		"April",
		"30th",
		"Jimmy",
		"St James",
	}
}
