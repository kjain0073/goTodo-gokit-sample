package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kjain0073/go-Todo/middlewareTodo"
	"github.com/kjain0073/go-Todo/router"
)

func main() {
	stopChan := make(chan os.Signal) // correct way to stop channel
	signal.Notify(stopChan, os.Interrupt)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", middlewareTodo.HomeHandler)
	r.Mount("/todo", router.TodoHandlers())

	srv := &http.Server{
		Addr:         middlewareTodo.Port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		log.Println("listening on port", middlewareTodo.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen:%s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server stopped Gracefully!")

}
