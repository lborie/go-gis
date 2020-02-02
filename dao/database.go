package dao

import (
	"errors"
)

var ErrNoTransaction = errors.New("cannot use this method without transaction")
var DB Database

// Database Interface
type Database interface {
	Select1() (bool, error)
	GeoJSON() (string, error)
}
