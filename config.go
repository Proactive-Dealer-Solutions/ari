package ari

import (
	"errors"
	"fmt"
	"strings"
)

// Config represents a transport to the asterisk
// config ARI resource.
type Config interface {

	// Get gets the reference to a config object
	Get(key *Key) ConfigHandle

	// Data gets the data for the config object
	Data(key *Key) (*ConfigData, error)

	// Update creates or updates the given tuples
	Update(key *Key, tuples []ConfigTuple) error

	// Delete deletes the dynamic configuration object.
	Delete(key *Key) error
}

// ConfigData contains the data for a given configuration object
type ConfigData struct {
	// Key is the cluster-unique identifier for this configuration
	Key *Key `json:"key"`

	Class string
	Type  string
	Name  string

	Fields []ConfigTuple
}

// ID returns the ID of the ConfigData structure
func (cd *ConfigData) ID() string {
	return fmt.Sprintf("%s/%s/%s", cd.Class, cd.Type, cd.Name)
}

// ConfigTuple is the key-value pair that defines a configuration entry
type ConfigTuple struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

// A ConfigHandle is a reference to a Config object
// on the asterisk service
type ConfigHandle interface {

	// ID returns the unique identifier for the config object
	ID() string

	// Data gets the current data for the config handle
	Data() (*ConfigData, error)

	// Update creates or updates the given config tuples
	Update(tuples []ConfigTuple) error

	// Delete deletes the dynamic configuration object
	Delete() error
}

// ParseConfigID parses the provided Config ID into its Class, Type, and ID components
func ParseConfigID(input string) (class, kind, id string, err error) {
	pieces := strings.Split(input, "/")
	if len(pieces) < 3 {
		err = errors.New("invalid input ID")
		return
	}

	return pieces[0], pieces[1], pieces[2], nil
}
