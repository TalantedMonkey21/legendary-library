package handler

import (
	"net/http"

	"gorm.io/gorm"
)

func NewRouter(c *gorm.DB) http.Handler {
	r := &Router{Db:c}
	mux := http.NewServeMux()
	mux.HandleFunc("/health", r.Health)
	mux.HandleFunc("/create", r.CreateNote)
	return mux
}
