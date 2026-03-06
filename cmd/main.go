package main

import (
	"log"
	"net/http"

	"github.com/TalantedMonkey21/GoLectures/internal/config"
	myhttp "github.com/TalantedMonkey21/GoLectures/internal/transport/http"
)


var nextId = 0


func main (){
	cfg := config.Load()
	mux := myhttp.NewRouter()
	log.Println(nextId)
	log.Println("Server starts on 8080")
	http.ListenAndServe(cfg.Port, mux)
}