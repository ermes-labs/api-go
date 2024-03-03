package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

// Struct that represents an Area.
type Area struct {
	Node
	// The sub-areas.
	Areas []Area `json:"areas,omitempty"`
}

// NewArea creates a new Area.
func NewArea(node Node, areas []Area) (*Area, error) {
	area := &Area{
		Node:  node,
		Areas: areas,
	}

	return area, CheckArea(*area, math.Inf(1), make(map[string]bool), make(map[string]bool))
}

// CheckArea checks the Area.
func CheckArea(area Area, maxDepth float64, nameMap map[string]bool, hostMap map[string]bool) error {
	// Check the node.
	if err := CheckNode(area.Node); err != nil {
		return err
	}

	// Checks that the depth is not exceeded.
	if maxDepth <= 0 {
		return ErrAreaMaxDepth
	}

	// Checks that the name is unique.
	if nameMap[area.AreaName] {
		return fmt.Errorf("%w: %s", ErrAreaNodeNameUnique, area.AreaName)
	} else {
		// Adds the name to the map.
		nameMap[area.AreaName] = true
	}

	// Checks that the host is unique.
	if hostMap[area.Host] {
		return fmt.Errorf("%w: %s", ErrHostUnique, area.Host)
	} else {
		// Adds the host to the map.
		hostMap[area.Host] = true
	}

	// Checks that the sub-areas are valid.
	for _, subArea := range area.Areas {
		if err := CheckArea(subArea, maxDepth-1, nameMap, hostMap); err != nil {
			return err
		}
	}

	return nil
}

// UnmarshalArea unmarshals the Area.
func UnmarshalArea(data []byte) (Area, error) {
	var r Area
	err := json.Unmarshal(data, &r)

	if err == nil {
		err = CheckArea(r, math.Inf(1), make(map[string]bool), make(map[string]bool))
	}

	return r, err
}

// MarshalArea marshals the Area.
func MarshalArea(area Area) ([]byte, error) {
	err := CheckArea(area, math.Inf(1), make(map[string]bool), make(map[string]bool))

	if err != nil {
		return nil, err
	}

	return json.Marshal(area)
}

// Errors.
var (
	ErrAreaMaxDepth       = errors.New("area max depth reached")
	ErrAreaNodeNameUnique = errors.New("area name must be unique")
	ErrHostUnique         = errors.New("host must be unique")
)
