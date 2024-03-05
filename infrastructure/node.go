package infrastructure

import (
	"encoding/json"
	"errors"
)

// Type that represents a map of resources and their values.
type Resources = map[string]float64

// Struct that represents the data of a node.
type Node struct {
	// The name of the node.
	AreaName string `json:"areaName"`
	// The host of the node.
	Host string `json:"host"`
	// The geo coordinates of the node.
	GeoCoordinates GeoCoordinates `json:"geoCoordinates,omitempty"`
	// The resources of the node.
	Resources Resources `json:"resources,omitempty"`
}

// NewNode creates a new node.
func NewNode(areaName string, host string, geoCoordinates GeoCoordinates) (*Node, error) {
	node := &Node{
		AreaName:       areaName,
		Host:           host,
		GeoCoordinates: geoCoordinates,
	}

	return node, CheckNode(*node)
}

// CheckNode checks the node.
func CheckNode(node Node) error {
	// Checks that the name is not empty.
	if node.AreaName == "" {
		return ErrAreaNameEmpty
	}

	// Checks that the host is not empty.
	if node.Host == "" {
		return ErrHostEmpty
	}

	// Checks that the latitude is in the range [-90, 90].
	if node.GeoCoordinates.Latitude < -90 || node.GeoCoordinates.Latitude > 90 {
		return ErrLatitudeOutOfRange
	}

	// Checks that the longitude is in the range [-180, 180].
	if node.GeoCoordinates.Longitude < -180 || node.GeoCoordinates.Longitude > 180 {
		return ErrLongitudeOutOfRange
	}

	// Checks that the resources are not negative.
	for _, v := range node.Resources {
		if v < 0 && v != -1 {
			return ErrResourceValueNegative
		}
	}

	return nil
}

// UnmarshalNode unmarshals the node.
func UnmarshalNode(data []byte) (Node, error) {
	var r Node
	err := json.Unmarshal(data, &r)

	if err == nil {
		err = CheckNode(r)
	}

	return r, err
}

// MarshalNode marshals the node.
func MarshalNode(node Node) ([]byte, error) {
	err := CheckNode(node)

	if err != nil {
		return nil, err
	}

	return json.Marshal(node)
}

// Errors.
var (
	ErrAreaNameEmpty         = errors.New("area name cannot be empty")
	ErrHostEmpty             = errors.New("host cannot be empty")
	ErrLatitudeOutOfRange    = errors.New("latitude out of range")
	ErrLongitudeOutOfRange   = errors.New("longitude out of range")
	ErrResourceValueNegative = errors.New("resource value cannot be negative")
)
