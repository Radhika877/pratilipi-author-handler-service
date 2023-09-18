package main

import (
	"author-handler-service/lib"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var config = lib.Config{}

func main() {
	var secretConfig string
	lib.GetSecrets(&secretConfig, &config)
	r := mux.NewRouter()

	//Health check end-point
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health check success!")
	})

	//Author APIs
	authorContentRouter := r.PathPrefix("/author").Subrouter()
	authorContentRouter.HandleFunc("/content", PublishContent).Methods("POST")
	http.Handle("/", r)

	//Queue (Consumer) API
	queueRouter := r.PathPrefix("/queue").Subrouter()
	queueRouter.HandleFunc("/consumer", HandleConsumer)

	//Premium Author Configuration API
	configRouter := r.PathPrefix("/config").Subrouter()
	configRouter.HandleFunc("/", ConfigModification).Methods("PATCH")

	//Initialise mongodb
	mongoClient := lib.InitialiseMongoDb(&config)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())

	//Disconnect mongodb client in case of error
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	//spin up server
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", config.Port),
	}

	fmt.Println("starting service at port " + config.Port)
	log.Fatal(server.ListenAndServe())
}
