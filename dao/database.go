package dao

var DB Database

// Database Interface
type Database interface {
	Select1() (bool, error)
	GeoJSONDepartements() (string, error)
	GeoJSONRegions() (string, error)
	GeoJSONCovid(float64, float64) (string, error)
	GeoJSONSNCF() (string, error)
	GeoJSONSNCFParRegions() (string, error)
	GeoJSONSNCFParDepartements() (string, error)
}
