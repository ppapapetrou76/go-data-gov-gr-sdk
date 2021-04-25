package traffic

import "time"

// List is a representation of a Data array.
type List []Data

// Data describes the data returned by the traffic API.
type Data struct {
	DeviceID     string    `json:"deviceid"`
	CountedCars  int       `json:"countedcars"`
	ProcessTime  time.Time `json:"appprocesstime"`
	RoadName     string    `json:"road_name"`
	RoadInfo     string    `json:"road_info"`
	AverageSpeed float64   `json:"average_speed"`
}

// FilterByRoadName filters a traffic List by the road name and returns a new list that contains only the filtered
// data.
func (l List) FilterByRoadName(roadName string) List {
	newList := List{}

	for _, d := range l {
		if d.RoadName == roadName {
			newList = append(newList, d)
		}
	}

	return newList
}
