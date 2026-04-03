package handler

import (
	"net/http"

	"github.com/TalantedMonkey21/GoLectures/internal/models"
	"github.com/TalantedMonkey21/GoLectures/internal/response"
)

func (rt *Router) DeleteNote (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WriteJSONError(w, http.StatusMethodNotAllowed, "Incorrect method")
		return
	}
	var note models.Note
	if err := rt.Db.Delete(&note, r.URL.Query().Get("id")).Error; err != nil {
		response.WriteJSONError(w, http.StatusNotFound, "No note")
		return
	}
	response.WriteJSONResponse(w, http.StatusOK, note)
}