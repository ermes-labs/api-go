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

// Returns the flatten list of areas.
func (a *Area) Flatten() []*Area {
	// Append this to the areas returned by the sub-areas.
	areas := []*Area{a}

	// Append the sub-areas.
	for _, subArea := range a.Areas {
		areas = append(areas, subArea.Flatten()...)
	}

	return areas
}

// String returns the string representation of the Area.
func (a *Area) String() string {
	// Return the json string representation.
	data, _ := json.Marshal(a)
	return string(data)
}

// NewArea creates a new Area.
func NewArea(node Node, areas []Area) (*Area, map[string]*Area, error) {
	area := &Area{
		Node:  node,
		Areas: areas,
	}

	areasMap := make(map[string]*Area)

	return area, areasMap, CheckArea(*area, math.Inf(1), areasMap, make(map[string]bool))
}

// CheckArea checks the Area.
func CheckArea(area Area, maxDepth float64, areasMap map[string]*Area, hostsMap map[string]bool) error {
	// Check the node.
	if err := CheckNode(area.Node); err != nil {
		return err
	}

	// Checks that the depth is not exceeded.
	if maxDepth <= 0 {
		return ErrAreaMaxDepth
	}

	// Checks that the name is unique.
	if _, ok := areasMap[area.AreaName]; ok {
		return fmt.Errorf("%w: %s", ErrAreaNodeNameUnique, area.AreaName)
	} else {
		// Adds the name to the map.
		areasMap[area.AreaName] = &area
	}

	// Checks that the host is unique.
	if _, ok := hostsMap[area.Host]; ok {
		return fmt.Errorf("%w: %s", ErrHostUnique, area.Host)
	} else {
		// Adds the host to the map.
		hostsMap[area.Host] = true
	}

	// Checks that the sub-areas are valid.
	for _, subArea := range area.Areas {
		if err := CheckArea(subArea, maxDepth-1, areasMap, hostsMap); err != nil {
			return err
		}
	}

	return nil
}

// UnmarshalArea unmarshals the Area.
func UnmarshalArea(data []byte) (*Area, error) {
	var r Area
	err := json.Unmarshal(data, &r)

	if err == nil {
		err = CheckArea(r, math.Inf(1), make(map[string]*Area), make(map[string]bool))
	}

	return &r, err
}

// MarshalArea marshals the Area.
func MarshalArea(area Area) ([]byte, error) {
	err := CheckArea(area, math.Inf(1), make(map[string]*Area), make(map[string]bool))

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
