package foodinspection

// Data represents the data returned by the food inspections API.
type Data struct {
	Year                   int     `json:"year"`
	Inspections            int     `json:"inspections"`
	Violations             int     `json:"violations"`
	ViolatingOrganizations float64 `json:"violating_organizations"`
	Penalties              float64 `json:"penalties"`
}

// List is a representation of a Data array.
type List []Data

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
