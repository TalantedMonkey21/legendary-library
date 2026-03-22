package http

import (
	"net/http"
	"strconv"
)

func DeleteLecturesHandler(w http.ResponseWriter, r *http.Request){
	lectureID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Nots = append(Nots[:lectureID], Nots[lectureID+1:]...)

	w.WriteHeader(http.StatusNoContent)
}
