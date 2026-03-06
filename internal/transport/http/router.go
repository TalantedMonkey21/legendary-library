package http

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	mux.Handle("GET /lectures", http.HandlerFunc(GetLecturesHandler))
	mux.Handle("POST /lectures/create", http.HandlerFunc(CreateLectureHandler))
	mux.Handle("PUT /lectures/update", http.HandlerFunc(UpdateLecturesHandler))
	mux.Handle("DELETE /lectures/delete", http.HandlerFunc(DeleteLecturesHandler))
	mux.Handle("GET /lectures/all", http.HandlerFunc(GetAllLecturesHandler))
	return mux
}
