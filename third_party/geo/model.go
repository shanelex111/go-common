package geo

type GeoCity struct {
	IP            string  `json:"ip"`
	ContinentCode string  `json:"continent_code"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	SubvisionCode string  `json:"subvision_code"`
	SubvisionName string  `json:"subvision_name"`
	CityName      string  `json:"city_name"`
	Postal        string  `json:"postal"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	TimeZone      string  `json:"time_zone"`
}
