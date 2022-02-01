package api

import (
	"github.com/gorilla/mux"
)

func Init_Server(router *mux.Router) {
	router.HandleFunc("/projects", getProjectHandler).Methods("GET")
	router.HandleFunc("/projects", postProjectHandler).Methods("POST")
	router.HandleFunc("/newdata", newDataHandler).Methods("GET")
	router.HandleFunc("/level0", postLevel0).Methods("POST")
	router.Use(loggingMiddleware)
}
