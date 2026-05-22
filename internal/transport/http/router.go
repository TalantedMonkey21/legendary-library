package httptransport

import (
	"net/http"

	"github.com/TalantedMonkey21/GoLectures/internal/middleware"
)

// TODO:
// GetByID
// Update
// Delete
func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /health", http.HandlerFunc(handler.Health))
	mux.Handle("POST /create", middleware.ContentTypeJSON(http.HandlerFunc(handler.CreateNote)))
	mux.Handle("GET /notes/{id}", http.HandlerFunc(handler.GetNote))
	mux.Handle("POST /update/{id}", middleware.ContentTypeJSON(http.HandlerFunc(handler.UpdateNote)))
	mux.Handle("DELETE /delete/{id}", http.HandlerFunc(handler.DeleteNote))

	return middleware.RecoverMiddleware(middleware.LoggingMiddleware(mux))
}
