package handler

import (
	"net/http"

	"github.com/TalantedMonkey21/GoLectures/internal/middleware"
	"gorm.io/gorm"
)

type Router struct {
	Db *gorm.DB
}

type Middleware func(http.Handler) (http.Handler)

func Chain(next http.Handler, m ...Middleware) (http.Handler) {
	wrapped := next
	for _, f := range m {
		wrapped = f(wrapped)
	}
	return wrapped
}

func NewRouter(c *gorm.DB) http.Handler {
	r := &Router{Db:c}
	mux := http.NewServeMux()

	mux.Handle("GET /health", http.HandlerFunc(r.Health))
	mux.Handle("POST /create", middleware.ContentTypeMiddleware(http.HandlerFunc(r.CreateNote)))
	mux.Handle("GET /notes", http.HandlerFunc(r.GetNotes))
	mux.Handle("GET /note/{id}", http.HandlerFunc(r.GetNote))
	mux.Handle("DELETE /delete/{id}", http.HandlerFunc(r.DeleteNote))
	// mux.HandleFunc("/health", r.Health)
	// mux.HandleFunc("/create", r.CreateNote)
	// mux.HandleFunc("/notes", r.GetNotes)
	// mux.HandleFunc("/note/{id}", r.GetNote)
	// mux.HandleFunc("/delete", r.DeleteNote)

	newMux := Chain(mux, middleware.LoggingMiddleware, middleware.RecoverMiddleware)
	return newMux
}
