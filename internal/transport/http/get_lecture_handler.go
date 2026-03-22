package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"github.com/TalantedMonkey21/GoLectures/internal/config"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Created_at time.Time `json:"created_at"`
}

var Nots = []Note{
	{Id: 0, Title: "Pushkin", Body: "russian writer", Created_at: time.Now()},
	{Id: 1, Title: "Tolstoy", Body: "russian writer", Created_at: time.Now()},
	{Id: 2, Title: "Dostoevsky", Body: "russian writer", Created_at: time.Now()},
}
func GetLecturesHandler(w http.ResponseWriter, r *http.Request){
	lectureID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || lectureID >= len(Nots) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Nots[lectureID])
}

// func GetAllLecturesHandler(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(Notes)
// }

func GetAllLecturesHandler(w http.ResponseWriter, r *http.Request){
	var notes []Note
	if err := config.DB.Find(&notes).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}