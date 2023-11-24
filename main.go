package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/itsjustvaal/blogaggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DBURL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	queries := database.New(db)

	apiCfg := apiConfig{
		DB: queries,
	}

	mainRouter := chi.NewRouter()
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// v1 Router
	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", handleGetReadiness)
	v1Router.Get("/err", handleGetErr)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUserByAPIKey))
	v1Router.Get("/feeds", apiCfg.handlerGetAllFeeds)
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetAllFeedFollows))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsGet))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))

	v1Router.Delete("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	// Mount other routers
	mainRouter.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(queries, collectionConcurrency, collectionInterval)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
