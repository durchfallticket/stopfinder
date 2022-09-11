package stopfinder

import "math"

type Net struct {
	Name  string `json:"name"`
	Zones []int `json:"zones"`
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func NewCoordinate(x float64, y float64) Coordinate {
	return Coordinate{((2 * math.Atan(math.Exp(((((12000000-y)/2.003750834e7)*180)*3.141592653589793)/180))) - 1.5707963267948966) * 57.29577951308232, (x / 2.003750834e7) * 180}
}

type Info struct {
	Coord             Coordinate `json:"coords"`
	Elevation         int16 `json:"elevation"`
	Id                string `json:"id"`
	IsTransferStation bool `json:"is_transfer_station"`
	Level             int16 `json:"level"`
	Name              string `json:"name"`
	Nets              []Net `json:"nets"`
	Omc               int `json:"omc"`
	ParentId          string `json:"parent_id"`
	ParentName        string `json:"parent_name"`
	NormalizedName    string `json:"-"`
}
