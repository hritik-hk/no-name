package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/hritik-hk/rss-aggregator/handlers"
	"github.com/hritik-hk/rss-aggregator/internal/service"
	"github.com/hritik-hk/rss-aggregator/utils"

	"github.com/hritik-hk/rss-aggregator/config"
)

func main() {

	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT was not found in environment")
	}

	DbConfig, err := config.NewDbConfig()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	//scraper configs
	const collectionConcurrency = 5
	const collectionInterval = time.Minute
	go service.StartScraping(DbConfig.DB, collectionConcurrency, collectionInterval)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//create handlers
	userHandler := handlers.UserHandler{DB: DbConfig.DB}
	feedHandler := handlers.FeedHandler{DB: DbConfig.DB}
	feedFollowsHandler := handlers.FeedFollowsHandler{DB: DbConfig.DB}

	v1Router := chi.NewRouter()
	v1Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, 201, "backend working")
	})

	v1Router.Get("/err", handlers.HandlerErr)

	v1Router.Post("/users", userHandler.CreateUser)
	v1Router.Get("/users", DbConfig.MiddlewareAuth(userHandler.GetUser))

	v1Router.Post("/feeds", DbConfig.MiddlewareAuth(feedHandler.CreateFeed))
	v1Router.Get("/feeds", feedHandler.GetFeeds)

	v1Router.Post("/feed_follows", DbConfig.MiddlewareAuth(feedFollowsHandler.CreateFeedFollows))
	v1Router.Get("/feed_follows", DbConfig.MiddlewareAuth(feedFollowsHandler.GetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", DbConfig.MiddlewareAuth(feedFollowsHandler.DeleteFeedFollow))

	router.Mount("/api/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("server starting at PORT: %v", PORT)

	serverErr := server.ListenAndServe()

	if serverErr != nil {
		log.Fatal(serverErr)
	}

}
