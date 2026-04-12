package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/TalantedMonkey21/GoLectures/internal/models"
	"github.com/TalantedMonkey21/GoLectures/internal/response"
)

func (rt *Router) DeleteNote (w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	pathId, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "Invalid id")
		return
	}
	if r.Method != http.MethodDelete {
		response.WriteJSONError(w, http.StatusMethodNotAllowed, "Incorrect method")
		return
	}
	var note models.Note
	result := rt.Db.Delete(&note, pathId)

	if result.Error != nil {
		response.WriteJSONError(w, http.StatusNotFound, "Not found")
		return
	}
	if result.RowsAffected == 0 {
		response.WriteJSONError(w, http.StatusNotFound, "Not found")
		return
	}
	response.WriteJSONResponse(w, http.StatusOK, note)
}