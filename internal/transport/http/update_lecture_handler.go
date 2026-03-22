package http

import (
	"net/http"
	"time"
	"encoding/json"
	"strconv"
)


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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingNote)
}