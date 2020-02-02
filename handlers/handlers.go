package handlers

import (
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