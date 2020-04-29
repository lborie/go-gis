package handlers

import (
	"fmt"
	"github.com/lborie/go-gis/dao"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Regions(w http.ResponseWriter, _ *http.Request) {
	logrus.Info("requesting Geojson Regions")
	requestGeojson(w, dao.DB.GeoJSONRegions)
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
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write([]byte(result))
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func requestGeojson(w http.ResponseWriter, f func()(string, error)){
	result, err := f()
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write([]byte(result))
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}