package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TalantedMonkey21/GoLectures/internal/config"
	myhttp "github.com/TalantedMonkey21/GoLectures/internal/transport/http"
)




func main (){
	cfg := config.Load()
	db := cfg.Db.GetDsn()
	fmt.Println(db)
	mux := myhttp.NewRouter()
	log.Println("Server starts on 8080")
	http.ListenAndServe(cfg.Port, mux)
}