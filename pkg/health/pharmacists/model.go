package pharmacist

// List is a representation of a Data array.
type List []Data

// Data describes the data returned by the pharmacists API.
type Data struct {
	Year     int    `json:"year"`
	Quarter  string `json:"quarter"`
	Active   int    `json:"active"`
	Entrants int   `json:"entrants"`
	Exits    int   `json:"exits"`
}

// FilterByYearRange filters a food inspection List based on the year range provided and returns a new list that
// contains only the filtered data.
func (l List) FilterByYearRange(fromYear, toYear int) List {
	newList := List{}

	for _, d := range l {
		if d.Year >= fromYear && d.Year <= toYear {
			newList = append(newList, d)
		}
	}

	return newList
}
