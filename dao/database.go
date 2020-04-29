package dao

var DB Database

// Database Interface
type Database interface {
	GeoJSONRegions() (string, error)
	GeoJSONCovid(float64, float64) (string, error)
}
