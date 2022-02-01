package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Violent-Idiot/Portfolio-backend/pkg/api"
	"github.com/Violent-Idiot/Portfolio-backend/pkg/db"
	"github.com/gorilla/mux"
)

func main() {
	log.Print("working")
	router := mux.NewRouter()
	dbConfig := &db.Config{
		MongoUri: "",
	}
	// coll := &db.Collections{}
	// log.Printf("%p", coll)
	dbConfig.Db_init()
	// log.Print(coll)
	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	api.Init_Server(router)
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Print("Shutting down the server")
	os.Exit(0)
}
