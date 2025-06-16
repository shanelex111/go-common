package geo

import "net"

func GetCity(ip string) (*GeoCity, error) {
	ipAddress := net.ParseIP(ip)
	record, err := DB.City(ipAddress)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, nil
	}

	result := &GeoCity{
		ContinentCode: record.Continent.Code,
		CountryCode:   record.Country.IsoCode,
		CountryName:   record.Country.Names["en"],
		CityName:      record.City.Names["en"],
		Postal:        record.Postal.Code,
		Latitude:      record.Location.Latitude,
		Longitude:     record.Location.Longitude,
		TimeZone:      record.Location.TimeZone,
	}
	if len(record.Subdivisions) > 0 {
		result.SubvisionCode = record.Subdivisions[0].IsoCode
		result.SubvisionName = record.Subdivisions[0].Names["en"]
	}

	return result, nil

}
