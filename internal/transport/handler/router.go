package handler

import (
	"net/http"

	"gorm.io/gorm"
)

type Router struct {
	Db *gorm.DB
}

func NewRouter(c *gorm.DB) http.Handler {
	r := &Router{Db:c}
	mux := http.NewServeMux()
	mux.HandleFunc("/health", r.Health)
	mux.HandleFunc("/create", r.CreateNote)
	mux.HandleFunc("/notes", r.GetNotes)
	mux.HandleFunc("/note", r.GetNote)
	mux.HandleFunc("/delete", r.DeleteNote)
	return mux
}
