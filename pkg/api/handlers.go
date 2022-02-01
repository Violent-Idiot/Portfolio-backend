package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Violent-Idiot/Portfolio-backend/pkg/db"
)

func getProjectHandler(w http.ResponseWriter, req *http.Request) {
	data := db.GetProjects()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
func newDataHandler(w http.ResponseWriter, req *http.Request) {

	data := db.GetProjects()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
func postProjectHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	data := r.FormValue("data")
	log.Println(name, data)
	db.PostProject(name, data)
	json.NewEncoder(w).Encode("posted")
}
func postLevel0(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	data := r.FormValue("data")
	dir := r.FormValue("dir")
	result := db.PostLevel0(name, data, dir)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
