package handler

import (
	"net/http"

	"github.com/TalantedMonkey21/GoLectures/internal/models"
	"github.com/TalantedMonkey21/GoLectures/internal/response"
)

func (rt *Router) GetNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteJSONError(w, http.StatusMethodNotAllowed, "Incorrect method")
		return
	}
	var notes []models.Note
	if err := rt.Db.Find(&notes).Error; err != nil {
		response.WriteJSONError(w, http.StatusNotFound, "No note")
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, notes)
}

func (rt *Router) GetNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteJSONError(w, http.StatusMethodNotAllowed, "Incorrect method")
		return
	}
	var note models.Note
	if err := rt.Db.First(&note, r.URL.Query().Get("id")).Error; err != nil {
		response.WriteJSONError(w, http.StatusNotFound, "No note")
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, note)
}