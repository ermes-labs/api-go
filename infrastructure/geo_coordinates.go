package infrastructure

import (
	"encoding/json"
	"errors"
)

type GeoCoordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// NewGeoCoordinates creates a new GeoCoordinates.
func NewGeoCoordinates(longitude float64, latitude float64) (*GeoCoordinates, error) {
	coordinates := &GeoCoordinates{
		Longitude: longitude,
		Latitude:  latitude,
	}

	return coordinates, CheckGeoCoordinates(*coordinates)
}

// CheckGeoCoordinates checks if the GeoCoordinates are valid.
func CheckGeoCoordinates(coordinates GeoCoordinates) error {
	if coordinates.Longitude < -180 || coordinates.Longitude > 180 {
		return ErrGeoCoordinatesLongitude
	}

	if coordinates.Latitude < -90 || coordinates.Latitude > 90 {
		return ErrGeoCoordinatesLatitude
	}

	return nil
}

// UnmarshalGeoCoordinates unmarshals the GeoCoordinates.
func UnmarshalGeoCoordinates(data []byte) (GeoCoordinates, error) {
	var r GeoCoordinates
	err := json.Unmarshal(data, &r)

	if err == nil {
		err = CheckGeoCoordinates(r)
	}

	return r, err
}

// MarshallGeoCoordinates marshalls the GeoCoordinates.
func MarshallGeoCoordinates(coordinates GeoCoordinates) ([]byte, error) {
	if err := CheckGeoCoordinates(coordinates); err != nil {
		return nil, err
	}

	return json.Marshal(coordinates)
}

var (
	// ErrGeoCoordinatesLongitude is returned when the longitude is not valid.
	ErrGeoCoordinatesLongitude = errors.New("longitude is not valid")
	// ErrGeoCoordinatesLatitude is returned when the latitude is not valid.
	ErrGeoCoordinatesLatitude = errors.New("latitude is not valid")
)
