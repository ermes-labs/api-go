package infrastructure_test

import (
	"errors"
	"testing"

	"github.com/ermes-labs/api-go/infrastructure"
)

func TestInvalidGeoCoordinates(t *testing.T) {
	nodeJson := `{
		"areaName": "area",
		"host": "host",
		"geoCoordinates": {
			"latitude": 91,
			"longitude": 181
		}
	}`

	_, err := infrastructure.UnmarshalNode([]byte(nodeJson))

	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if !errors.Is(err, infrastructure.ErrLatitudeOutOfRange) {
		t.Errorf("Expected error %v, got %v", infrastructure.ErrLatitudeOutOfRange, err)
	}
}
