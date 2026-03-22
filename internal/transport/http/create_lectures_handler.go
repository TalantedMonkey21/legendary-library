package http

import (
	"net/http"
	"log"
	"time"
	"encoding/json"
)

type InputNote struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

func CreateLectureHandler(w http.ResponseWriter, r *http.Request){
	newId := len(Nots)
	var note InputNote
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(note)
	if note.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resultNote := Note{
		Id: newId,
		Title: note.Title,
		Body: note.Body,
		Created_at: time.Now(),
	}
	Nots = append(Nots, resultNote)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultNote)
}