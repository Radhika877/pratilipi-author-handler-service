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
	authorContentRouter.HandleFunc("/content", PublishContent).Methods("POST")                      //API for author to publish content followed by premium status ops
	authorContentRouter.HandleFunc("/{authorId}/followers", SimulateAuthorFollowers).Methods("PUT") // API to simulate author's followers count
	authorContentRouter.HandleFunc("/sync-premium", PremiumAuthorDailySync).Methods("POST")         //API to hit daily sync funcion to update author's premium flag
	http.Handle("/", r)

	//Queue (Consumer) API
	queueRouter := r.PathPrefix("/queue").Subrouter()
	queueRouter.HandleFunc("/consumer", HandleConsumer).Methods("POST") //API being used as webhook to handle consumer

	//Premium Author Configuration API
	configRouter := r.PathPrefix("/config").Subrouter()
	configRouter.HandleFunc("/", ConfigModification).Methods("PATCH") //API to alter global config for premium author eligiblity

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
