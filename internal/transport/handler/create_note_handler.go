package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TalantedMonkey21/GoLectures/internal/models"
	"github.com/TalantedMonkey21/GoLectures/internal/response"
)

func (rt *Router) CreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteJSONError(w, http.StatusMethodNotAllowed, "Incorrect method")
		return
	}
	var n models.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if len(n.Content) < 3 {
		response.WriteJSONError(w, http.StatusBadRequest, "Write more!!!")
		return
	}
	if err := rt.Db.Create(&n).Error; err != nil {
		response.WriteJSONError(w, http.StatusInternalServerError, "Cannot create note")
		return
	}
	response.WriteJSONResponse(w, http.StatusCreated, n)
}