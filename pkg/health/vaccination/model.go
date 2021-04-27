package vaccination

// List is a representation of a Data array.
type List []Data

// Data describes the data returned by the vaccination API.
type Data struct {
	Area              string `json:"area"`
	AreaID            int    `json:"areaid"`
	DailyShot1        int    `json:"dailydose1"`
	DailyShot2        int    `json:"dailydose2"`
	DayDiff           int    `json:"daydiff"`
	DayTotal          int    `json:"daytotal"`
	TotalPersons      int    `json:"totaldistinctpersons"`
	TotalShot1        int    `json:"totaldose1"`
	TotalShot2        int    `json:"totaldose2"`
	TotalVaccinations int    `json:"totalvaccinations"`
	ReferenceDate     string `json:"referencedate"`
}

// FilterByArea filters a vaccination List by the given area name and returns a new list that contains only the filtered
// data.
func (l List) FilterByArea(areaName string) List {
	if areaName == "" {
		return l
	}
	newList := List{}

	for _, d := range l {
		if d.Area == areaName {
			newList = append(newList, d)
		}
	}

	return newList
}
