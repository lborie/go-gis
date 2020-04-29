package handlers

import (
	"fmt"
	"github.com/lborie/go-gis/dao"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func RenderMap(w http.ResponseWriter, _ *http.Request) {
	googleMapsKey := os.Getenv("GOOGLE_MAPS_KEY")
	if googleMapsKey == "" {
		logrus.Error("missing google maps key")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := template.Must(template.ParseFiles(filepath.Join("templates", "index.tmpl")))
	_ = t.Execute(w, map[string]interface{}{
		"GOOGLE_MAPS_KEY": googleMapsKey,
	})
}

func Regions(w http.ResponseWriter, _ *http.Request) {
	logrus.Info("requesting Geojson Regions")
	requestGeojson(w, dao.DB.GeoJSONRegions)
}

func Departements(w http.ResponseWriter, _ *http.Request) {
	logrus.Info("requesting Geojson Departements")
	requestGeojson(w, dao.DB.GeoJSONDepartements)
}

func Covid(w http.ResponseWriter, r *http.Request) {
	logrus.Info("requesting Geojson Covid")
	latString := r.URL.Query().Get("lat")
	lat, err := strconv.ParseFloat(latString, 64)
	if err != nil {
		logrus.Errorf(fmt.Errorf("lat parsing error : %w", err).Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	longString := r.URL.Query().Get("long")
	long, err := strconv.ParseFloat(longString, 64)
	if err != nil {
		logrus.Errorf(fmt.Errorf("long parsing error : %w", err).Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := dao.DB.GeoJSONCovid(lat, long)
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(result))
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func SNCF(w http.ResponseWriter, _ *http.Request) {
	logrus.Info("requesting Geojson SNCF")
	requestGeojson(w, dao.DB.GeoJSONSNCF)
}

func SNCFParRegions(w http.ResponseWriter, _ *http.Request) {
	logrus.Info("requesting Geojson SNCF Par Regions")
	requestGeojson(w, dao.DB.GeoJSONSNCFParRegions)
}

func SNCFParDepartements(w http.ResponseWriter, _ *http.Request) {
	logrus.Info("requesting Geojson SNCF Par Departements")
	requestGeojson(w, dao.DB.GeoJSONSNCFParDepartements)
}

func requestGeojson(w http.ResponseWriter, f func()(string, error)){
	result, err := f()
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(result))
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}