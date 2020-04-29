package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	maxIdleConns    = 5
	maxOpenConns    = 5
	connMaxLifetime = 5 * time.Minute
)

type databasePostgreSQL struct {
	session *sql.DB
}

// NewDatabasePostgreSQL returns a new dao with postgres cnx
func InitDatabasePostgreSQL(connectionURI string) (Database, error) {
	logrus.WithField("function", "NewDatabasePostgreSQL").WithField("connectionURI", connectionURI).Debug()

	db, err := sql.Open("postgres", connectionURI)

	if err != nil {
		return nil, fmt.Errorf("unable to get a connection to the postgres db: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to ping the postgres db: %v", err)
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	return &databasePostgreSQL{session: db}, nil
}

// STEP 1
func (db *databasePostgreSQL) GeoJSONRegions() (string, error) {
	var result string

	if err := db.session.QueryRow(`
			select json_build_object(
				'type', 'FeatureCollection',
				'features', json_agg(ST_AsGeoJSON(r.*)::json)
				)
			from "regions-20180101" r
			where r.nom = 'Hauts-de-France'
`).Scan(&result); err != nil {
		return "", err
	}
	return result, nil
}

func (db *databasePostgreSQL) GeoJSONCovid(lat, long float64) (string, error) {
	var result string

	query := fmt.Sprintf(`
select json_build_object(
	   'type', 'FeatureCollection',
	   'features', json_agg(
			   ST_AsGeoJSON(st_intersection(st_transform(st_buffer(st_transform(st_setsrid(st_makepoint(%f, %f), 4326), 2154), 100000), 4326), france.geom))::json
		   )
   )
from france
`, long, lat)

	if err := db.session.QueryRow(query).Scan(&result); err != nil {
		return "", err
	}
	return result, nil
}
