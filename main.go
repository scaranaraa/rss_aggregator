package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/scaranaraa/rss_aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main(){

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL not set")
	}
	if port == "" {
		log.Fatal("Port not set")
	}

	conn, err1  := sql.Open("postgres",dbURL)
	if err1 != nil {
		
		log.Fatal("Cant connect to database: ",err1)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Post("/users",apiCfg.handlerCreateUser)
	v1Router.Get("/users",apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Get("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds/get", apiCfg.handlerGetFeeds)

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
