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
	maxOpenConns    = 10
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

func (db *databasePostgreSQL) Select1() (bool, error) {
	var result bool
	if err := db.session.QueryRow("select 1").Scan(&result); err != nil {
		return false, err
	}
	return result, nil
}

func (db *databasePostgreSQL) GeoJSONDepartements() (string, error) {
	var result string

	if err := db.session.QueryRow(`
			select json_build_object(
				'type', 'FeatureCollection',
				'features', json_agg(ST_AsGeoJSON(d.*)::json)
				)
			from "departements-20180101" d
			join "regions-20180101" r on st_within(d.geom, r.geom)
			where r.nom = 'Hauts-de-France'
`).Scan(&result); err != nil {
		return "", err
	}
	return result, nil
}

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

func (db *databasePostgreSQL) GeoJSONSNCF() (string, error) {
	var result string

	if err := db.session.QueryRow(`
			select json_build_object(
				'type', 'FeatureCollection',
				'features', json_agg(ST_AsGeoJSON(sncf.*)::json)
				)
			from "formes-des-lignes-du-rfn" sncf
`).Scan(&result); err != nil {
		return "", err
	}
	return result, nil
}