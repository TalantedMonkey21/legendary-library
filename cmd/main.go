package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TalantedMonkey21/GoLectures/internal/config"
	"github.com/TalantedMonkey21/GoLectures/internal/transport/handler"
)




func main (){
	cfg := config.Load()

	connect, err := config.ConnectDb(cfg.Db)
	if err != nil {
		log.Println("Cannot connect tot database", err)
		os.Exit(1)
	}
	log.Println("Migrate complete!")
	mux := handler.NewRouter(connect)
	http.ListenAndServe(cfg.Port, mux)
}