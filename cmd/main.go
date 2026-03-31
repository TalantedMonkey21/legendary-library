package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TalantedMonkey21/GoLectures/internal/config"
	"github.com/TalantedMonkey21/GoLectures/internal/db"
	"github.com/TalantedMonkey21/GoLectures/internal/transport/handler"
)




func main (){
	cfg := config.Load()

	connect, err := db.ConnectDb(cfg.Db)
	if err != nil {
		fmt.Println("Cannot connect tot database", err)
		os.Exit(1)
	}
	fmt.Println("Migrate complete!")
	mux := handler.NewRouter(connect)
	http.ListenAndServe(cfg.Port, mux)
}