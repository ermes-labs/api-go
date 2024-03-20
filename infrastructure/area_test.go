package infrastructure_test

import (
	"errors"
	"testing"

	"github.com/ermes-labs/api-go/infrastructure"
)

func TestInvalidNonUniqueAreasNames(t *testing.T) {
	areaJson := `{
		"areaName": "area",
		"host": "host",
		"areas": [
			{
				"areaName": "area1",
				"host": "host1"
			},
			{
				"areaName": "area2",
				"host": "host2",
				"areas": [
					{
						"areaName": "area1",
						"host": "host3"
					}
				]
			}
		]
	}`

	_, err := infrastructure.UnmarshalArea([]byte(areaJson))

	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if !errors.Is(err, infrastructure.ErrAreaNodeNameUnique) {
		t.Errorf("Expected error %v, got %v", infrastructure.ErrAreaNodeNameUnique, err)
	}
}

func TestInvalidNonUniqueHosts(t *testing.T) {
	areaJson := `{
		"areaName": "area",
		"host": "host",
		"areas": [
			{
				"areaName": "area1",
				"host": "host1"
			},
			{
				"areaName": "area2",
				"host": "host2",
				"areas": [
					{
						"areaName": "area3",
						"host": "host1"
					}
				]
			}
		]
	}`

	_, err := infrastructure.UnmarshalArea([]byte(areaJson))

	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if !errors.Is(err, infrastructure.ErrHostUnique) {
		t.Errorf("Expected error %v, got %v", infrastructure.ErrHostUnique, err)
	}
}
