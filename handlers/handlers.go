package handlers

import (
	"fmt"
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

func GeoJson(w http.ResponseWriter, r *http.Request) {
	select1, err := dao.DB.Select1()
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	logrus.Info(select1)
	_, err = w.Write([]byte(fmt.Sprintf("voilà le résultat : %v", select1)))
	if err != nil {
		logrus.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}