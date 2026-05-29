package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TalantedMonkey21/GoLectures/internal/config"
	"github.com/TalantedMonkey21/GoLectures/internal/repository"
	httptransport "github.com/TalantedMonkey21/GoLectures/internal/transport/http"
	"github.com/TalantedMonkey21/GoLectures/internal/usecase"
	"gorm.io/gorm"
)


func doRetryConnection(cfg *config.Config) (*gorm.DB, error) {
	maxAttempts := 10
	for i := range(maxAttempts) {
		db, err := config.ConnectDB(cfg.DB)
		if err == nil {
			return db, nil
		}
		log.Printf("Error connecting to database: %v, retrying in 3 seconds...\n", err)
		if i < maxAttempts {
			time.Sleep(3 * time.Second)
		}
	}
	return nil, errors.New("failed to connect to database after 10 attempts")
}

func main() {
	// получаем конфиги
	cfg := config.Load()

	// TODO:
	// может случится такое, что сервер поднимается раньше через БД
	// и сейчас подключение происходит 1 раз
	// нужно подумать над тем как сделать 10 попыток каждые 3 секунды

	// создаем соединение с бд
	db, err := doRetryConnection(cfg)
	if err != nil {
		log.Println("can not connect to database:", err)
		os.Exit(1)
	}

	// автомиграция
	if err := config.AutoMigrate(db); err != nil {
		log.Println("can not migrate to database:", err)
		os.Exit(1)
	}

	// слой хранилища
	noteRepo := repository.NewNoteRepo(db)

	// слой бизнес логики
	noteService := usecase.New(noteRepo)

	// слой хэндлеров
	handler := httptransport.NewHandler(noteService)

	// роутер
	router := httptransport.NewRouter(handler)

	// TODO
	// сделай Gracefull Shutdown
	// необязательно! если что то обсудим на уроке
	http.ListenAndServe(cfg.Port, router)
}
