package handlers

import (
	"github.com/lborie/go-gis/dao"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func RenderMap(w http.ResponseWriter, r *http.Request) {
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

func Regions(w http.ResponseWriter, r *http.Request) {
	logrus.Info("requesting Geojson Regions")
	result, err := dao.DB.GeoJSONRegions()
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write([]byte(result))
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Departements(w http.ResponseWriter, r *http.Request) {
	logrus.Info("requesting Geojson Departements")
	result, err := dao.DB.GeoJSONDepartements()
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write([]byte(result))
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}