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

// NewInfrastructure creates a new Infrastructure.
func NewInfrastructure(areaIdentifiers []string, areas []Area) (*Infrastructure, error) {
	infrastructure := &Infrastructure{
		AreaIdentifiers: areaIdentifiers,
		Areas:           areas,
	}

	return infrastructure, CheckInfrastructure(*infrastructure)
}

// CheckInfrastructure checks the Infrastructure.
func CheckInfrastructure(infrastructure Infrastructure) error {
	// Checks that the area identifiers are not empty.
	if len(infrastructure.AreaIdentifiers) == 0 {
		return ErrInfrastructureAreaIdentifiersEmpty
	}

	// Checks that all the identifiers are unique and not empty.
	identifiers := make(map[string]bool)
	for _, identifier := range infrastructure.AreaIdentifiers {
		// Checks that the identifier is not empty.
		if identifier == "" {
			return ErrInfrastructureAreaIdentifierEmpty
		}

		// Checks that the identifier is not duplicated.
		if _, ok := identifiers[identifier]; ok {
			return fmt.Errorf("%w: %s", ErrInfrastructureAreaIdentifiersUnique, identifier)
		}

		// Adds the identifier to the map.
		identifiers[identifier] = true
	}

	nameMap := make(map[string]bool)
	hostMap := make(map[string]bool)
	maxDepth := float64(len((infrastructure.AreaIdentifiers)))
	// Checks that all the areas are valid.
	for _, area := range infrastructure.Areas {
		if err := CheckArea(area, maxDepth, nameMap, hostMap); err != nil {
			return err
		}
	}

	return nil
}

// UnmarshalInfrastructure unmarshals the Infrastructure.
func UnmarshalInfrastructure(data []byte) (Infrastructure, error) {
	var r Infrastructure
	err := json.Unmarshal(data, &r)

	if err == nil {
		err = CheckInfrastructure(r)
	}

	return r, err
}

// MarshalInfrastructure marshals the Infrastructure.
func MarshalInfrastructure(infrastructure Infrastructure) ([]byte, error) {
	if err := CheckInfrastructure(infrastructure); err != nil {
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
