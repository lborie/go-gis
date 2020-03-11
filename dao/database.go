package dao

var DB Database

// Database Interface
type Database interface {
	Select1() (bool, error)
	GeoJSONDepartements() (string, error)
	GeoJSONRegions() (string, error)
	GeoJSONSNCF() (string, error)
	GeoJSONSNCFParRegions() (string, error)
	GeoJSONSNCFParDepartements() (string, error)
}
