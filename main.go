package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz",handlerReadiness)
	router.Mount("/v1",v1Router)
	v1Router.Get("/err", handleErr)
	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Listening on port %s\n", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
