package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Struct that represents the entry point of the infrastructure.
type Infrastructure struct {
	// Area identifiers are used to identify the hierarchy of areas.
	AreaIdentifiers []string `json:"areaIdentifiers"`
	// Areas are the hierarchy of areas.
	Areas []Area `json:"areas"`
}

// Returns the flatten list of areas.
func (i *Infrastructure) Flatten() []*Area {
	// Append the sub-areas.
	areas := make([]*Area, 0, len(i.Areas))
	for _, subArea := range i.Areas {
		areas = append(areas, subArea.Flatten()...)
	}

	return areas
}

// NewInfrastructure creates a new Infrastructure.
func NewInfrastructure(areaIdentifiers []string, areas []Area) (*Infrastructure, map[string]*Area, error) {
	infrastructure := &Infrastructure{
		AreaIdentifiers: areaIdentifiers,
		Areas:           areas,
	}

	areasMap, err := CheckInfrastructure(*infrastructure)

	return infrastructure, areasMap, err
}

// CheckInfrastructure checks the Infrastructure.
func CheckInfrastructure(infrastructure Infrastructure) (map[string]*Area, error) {
	// Checks that the area identifiers are not empty.
	if len(infrastructure.AreaIdentifiers) == 0 {
		return nil, ErrInfrastructureAreaIdentifiersEmpty
	}

	// Checks that all the identifiers are unique and not empty.
	identifiers := make(map[string]bool)
	for _, identifier := range infrastructure.AreaIdentifiers {
		// Checks that the identifier is not empty.
		if identifier == "" {
			return nil, ErrInfrastructureAreaIdentifierEmpty
		}

		// Checks that the identifier is not duplicated.
		if _, ok := identifiers[identifier]; ok {
			return nil, fmt.Errorf("%w: %s", ErrInfrastructureAreaIdentifiersUnique, identifier)
		}

		// Adds the identifier to the map.
		identifiers[identifier] = true
	}

	areasMap := make(map[string]*Area)
	hostsMap := make(map[string]bool)
	maxDepth := float64(len((infrastructure.AreaIdentifiers)))
	// Checks that all the areas are valid.
	for _, area := range infrastructure.Areas {
		if err := CheckArea(area, maxDepth, areasMap, hostsMap); err != nil {
			return nil, err
		}
	}

	return areasMap, nil
}

// UnmarshalInfrastructure unmarshals the Infrastructure.
func UnmarshalInfrastructure(data []byte) (*Infrastructure, map[string]*Area, error) {
	var r Infrastructure
	err := json.Unmarshal(data, &r)

	if err == nil {
		areasMap, err := CheckInfrastructure(r)

		if err == nil {
			return &r, areasMap, nil
		}
	}

	return &r, nil, err
}

// MarshalInfrastructure marshals the Infrastructure.
func MarshalInfrastructure(infrastructure Infrastructure) ([]byte, error) {
	if _, err := CheckInfrastructure(infrastructure); err != nil {
		return nil, err
	}

	return json.Marshal(infrastructure)
}

// Errors.
var (
	// ErrInfrastructureAreaIdentifiersUnique is returned when an node identifier is duplicated.
	ErrInfrastructureAreaIdentifiersUnique = errors.New("infrastructure node identifiers unique")
	// ErrInfrastructureAreaIdentifiersEmpty is returned when the Infrastructure has no node identifiers.
	ErrInfrastructureAreaIdentifiersEmpty = errors.New("infrastructure node identifiers empty")
	// ErrInfrastructureNodeIdentifierEmpty is returned when an node identifier is empty.
	ErrInfrastructureAreaIdentifierEmpty = errors.New("infrastructure node identifier empty")
)
