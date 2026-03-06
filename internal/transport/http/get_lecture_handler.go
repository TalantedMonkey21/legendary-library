package http

import (
	"encoding/json"
	"net/http"
	"time"
	"strconv"
)

type Note struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Created_at time.Time `json:"created_at"`
}
var Notes = []Note{
	{Id: 0, Title: "Pushkin", Body: "russian writer", Created_at: time.Now()},
	{Id: 1, Title: "Tolstoy", Body: "russian writer", Created_at: time.Now()},
	{Id: 2, Title: "Dostoevsky", Body: "russian writer", Created_at: time.Now()},
}
func GetLecturesHandler(w http.ResponseWriter, r *http.Request){
	lectureID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Notes[lectureID])
	w.WriteHeader(http.StatusOK)
}

func CreateLectureHandler(w http.ResponseWriter, r *http.Request){
	newId := len(Notes)
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if note.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	note.Id = newId
	Notes = append(Notes, note)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
	w.WriteHeader(http.StatusCreated)
}

func UpdateLecturesHandler(w http.ResponseWriter, r *http.Request){
	lectureID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var existingNote Note
	if existingNote.Id != lectureID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedNote Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if updatedNote.Title != "" {
		existingNote.Title = updatedNote.Title
	}
	if updatedNote.Body != "" {
		existingNote.Body = updatedNote.Body
	}

	existingNote.Created_at = time.Now()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingNote)
	w.WriteHeader(http.StatusOK)
}

func DeleteLecturesHandler(w http.ResponseWriter, r *http.Request){
	lectureID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Notes = append(Notes[:lectureID], Notes[lectureID+1:]...)

	w.WriteHeader(http.StatusNoContent)
}

func GetAllLecturesHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Notes)
	w.WriteHeader(http.StatusOK)
}