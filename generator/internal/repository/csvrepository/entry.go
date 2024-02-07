package csvrepository

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
	Child_1_First_Name       string
	Child_1_Saint_Name       string
	Child_1_Birthday_Month   string
	Child_1_Birthday_Day     string
	Child_2_First_Name       string
	Child_2_Saint_Name       string
	Child_2_Birthday_Month   string
	Child_2_Birthday_Day     string
	Child_3_First_Name       string
	Child_3_Saint_Name       string
	Child_3_Birthday_Month   string
	Child_3_Birthday_Day     string
	Child_4_First_Name       string
	Child_4_Saint_Name       string
	Child_4_Birthday_Month   string
	Child_4_Birthday_Day     string
	Child_5_First_Name       string
	Child_5_Saint_Name       string
	Child_5_Birthday_Month   string
	Child_5_Birthday_Day     string
	Child_6_First_Name       string
	Child_6_Saint_Name       string
	Child_6_Birthday_Month   string
	Child_6_Birthday_Day     string
	Child_7_First_Name       string
	Child_7_Saint_Name       string
	Child_7_Birthday_Month   string
	Child_7_Birthday_Day     string
}

// helper to convert a line from a csv to an entry
func parseLine(line []string) Entry {
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
			e.Child_1_First_Name = v
		case 21:
			e.Child_1_Saint_Name = v
		case 22:
			e.Child_1_Birthday_Month = v
		case 23:
			e.Child_1_Birthday_Day = v
		case 24:
			e.Child_2_First_Name = v
		case 25:
			e.Child_2_Saint_Name = v
		case 26:
			e.Child_2_Birthday_Month = v
		case 27:
			e.Child_2_Birthday_Day = v
		case 28:
			e.Child_3_First_Name = v
		case 29:
			e.Child_3_Saint_Name = v
		case 30:
			e.Child_3_Birthday_Month = v
		case 31:
			e.Child_3_Birthday_Day = v
		case 32:
			e.Child_4_First_Name = v
		case 33:
			e.Child_4_Saint_Name = v
		case 34:
			e.Child_4_Birthday_Month = v
		case 35:
			e.Child_4_Birthday_Day = v
		case 36:
			e.Child_5_First_Name = v
		case 37:
			e.Child_5_Saint_Name = v
		case 38:
			e.Child_5_Birthday_Month = v
		case 39:
			e.Child_5_Birthday_Day = v
		case 40:
			e.Child_6_First_Name = v
		case 41:
			e.Child_6_Saint_Name = v
		case 42:
			e.Child_6_Birthday_Month = v
		case 43:
			e.Child_6_Birthday_Day = v
		case 44:
			e.Child_7_First_Name = v
		case 45:
			e.Child_7_Saint_Name = v
		case 46:
			e.Child_7_Birthday_Month = v
		case 47:
			e.Child_7_Birthday_Day = v
		}
	}

	return e
}
