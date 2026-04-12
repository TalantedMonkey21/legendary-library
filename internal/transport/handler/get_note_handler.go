package handler

import (
	"net/http"
	"strconv"
	"strings"

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
	path := strings.Split(r.URL.Path, "/")
	pathId, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "Invalid id")
		return
	}
	if r.Method != http.MethodGet {
		response.WriteJSONError(w, http.StatusMethodNotAllowed, "Incorrect method")
		return
	}
	var note models.Note
	if err := rt.Db.First(&note, pathId).Error; err != nil {
		response.WriteJSONError(w, http.StatusNotFound, "Not found")
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, note)
}