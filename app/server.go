package app

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Server struct {
	db   *gorm.DB
	http *http.Server
}

func NewServer(db *gorm.DB) *Server {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Get("/article", deliveryPostApp(db).Articles)
	router.Post("/article", deliveryPostApp(db).CreatePost)
	router.Get("/article/{id}", deliveryPostApp(db).FindPostDetail)
	router.Put("/article/{id}", deliveryPostApp(db).UpddatePost)
	router.Delete("/article/{id}", deliveryPostApp(db).DeletePost)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              viper.GetString("listen_address"),
		Handler:           router,
	}

	server.SetKeepAlivesEnabled(true)

	srv := &Server{
		db:   db,
		http: server,
	}

	return srv
}

func (s *Server) Start(ctx context.Context) (err error) {
	fmt.Println("App: Starting")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err = s.http.ListenAndServe(); err != nil {
			return
		}
	}()

	<-ctx.Done()

	return
}

func (s *Server) Stop(ctx context.Context) (err error) {
	fmt.Println("App: Stopping...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err = s.http.Shutdown(ctx); err != nil {
			return
		}
	}()
	wg.Wait()

	return
}
